package args

import (
	"encoding/json"
	"errors"
	"os"
	"solidlabtest/src"
)

func ReadRawTree() (src.CommentsTree, error) {
	args := os.Args
	if len(args) <= 1 {
		return src.CommentsTree{}, errors.New("invalid args")
	}
	rawTree := args[1]

	var tree src.CommentsTree

	err := json.Unmarshal([]byte(rawTree), &tree)
	if err != nil {
		return src.CommentsTree{}, errors.New("invalid tree")
	}

	return tree, nil
}
