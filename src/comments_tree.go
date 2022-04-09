package src

import (
	"encoding/json"
	"net/http"
	er "solidlabtest/src/errors"
	"strconv"
	"sync"
)

type CommentsTree struct {
	Id      int            `json:"id,required"`
	Body    string         `json:"body"`
	Replies []CommentsTree `json:"replies"`
}

const URL = "https://25.ms/posts/"
const NUM_PARALLEL = 20

func (t *CommentsTree) SetBodies() error {
	bodies, err := t.getBodies()
	if err != nil {
		return err
	}

	t.setBodies(bodies)

	return nil
}

func (t *CommentsTree) setBodies(bodies map[int]string) {
	t.Body = bodies[t.Id]
	if t.Replies != nil {
		for i, reply := range t.Replies {
			reply.setBodies(bodies)
			t.Replies[i] = reply
		}
	}
}

// Stream inputs to input channel
func streamInputs(done <-chan struct{}, inputs []int) <-chan int {
	inputCh := make(chan int)
	go func() {
		defer close(inputCh)
		for _, input := range inputs {
			select {
			case inputCh <- input:
			case <-done:
				break
			}
		}
	}()
	return inputCh
}

type result struct {
	id      int
	bodyStr string
	err     error
}

func (t *CommentsTree) getIdsList(ids *[]int) {
	*ids = append(*ids, t.Id)
	if t.Replies != nil {
		for _, reply := range t.Replies {
			if reply.Replies != nil {
				reply.getIdsList(ids)
			}
		}
	}
}

func (t *CommentsTree) getBodies() (map[int]string, error) {
	var ids []int
	t.getIdsList(&ids)

	done := make(chan struct{})
	defer close(done)

	inputCh := streamInputs(done, ids)

	var wg sync.WaitGroup
	wg.Add(NUM_PARALLEL)

	resultCh := make(chan result)

	for i := 0; i < NUM_PARALLEL; i++ {
		// spawn N worker goroutines, each is consuming a shared input channel.
		go func() {
			for input := range inputCh {
				r := getBody(input)
				resultCh <- r
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := make(map[int]string)
	for el := range resultCh {
		if el.err != nil {
			return nil, el.err
		}
		results[el.id] = el.bodyStr
	}

	return results, nil
}

func getBody(id int) result {
	url := URL + strconv.Itoa(id)
	resp, err := http.Get(url)
	switch {
	case err != nil:
		return result{id, "", err}
	case resp.StatusCode == http.StatusNotFound:
		return result{id, "", er.ErrCommentNotFound}
	}

	defer resp.Body.Close()

	var respBody ServerResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respBody)
	if err != nil {
		return result{id, "", er.ErrInternal}
	}

	return result{id, respBody.Body, nil}
}
