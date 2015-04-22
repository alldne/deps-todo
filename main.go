package main

import (
	"fmt"
	"io/ioutil"

	"./lexer"
	"./parser"
	"./print"
	"./query"
)

func main() {
	dat, err := ioutil.ReadFile("sample.dtodo")
	if err != nil {
		panic(err)
	}

	todo(string(dat))
}

func todo(src string) {
	tokenChan := make(chan lexer.Token)
	treeChan := make(chan parser.Todo)
	go lexer.Run(tokenChan, &src)
	go parser.Run(treeChan, tokenChan)
	root := <-treeChan
	fmt.Printf("%s\n", print.Stringify(root))

	query.New(root)
	return
}
