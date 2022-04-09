package main

import (
	"encoding/json"
	"fmt"
	"solidlabtest/src/args"
)

func main() {
	rawTree, err := args.ReadRawTree()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = rawTree.SetBodies()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonTree, _ := json.Marshal(rawTree)

	fmt.Println(string(jsonTree))
}
