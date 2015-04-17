package main

import (
	"fmt"
	"io/ioutil"

	"./lexer"
)

func main() {
	dat, err := ioutil.ReadFile("sample.wtodo")
	if err != nil {
		panic(err)
	}

	todo(string(dat))
}

func todo(src string) {
	tokenChan := make(chan lexer.Token)
	go lexer.Run(tokenChan, &src)
	for t := range tokenChan {
		fmt.Printf("Got %s\n", t)
	}
	return
}
