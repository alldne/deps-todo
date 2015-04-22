package main

import (
	"fmt"
	"io/ioutil"

	"./lexer"
	"./parser"
	"./print"
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
	nodeChan := make(chan parser.Todo)
	go lexer.Run(tokenChan, &src)
	go parser.Run(nodeChan, tokenChan)
	root := <-nodeChan
	fmt.Printf("%s\n", print.Stringify(root))
	return
}
