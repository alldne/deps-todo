package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("sample.wtodo")
	if err != nil {
		panic(err)
	}

	todo(string(dat))
}

func todo(src string) {
	tokenChan := make(chan token)
	go lexer(tokenChan, src)
	for t := range tokenChan {
		fmt.Printf("Got token %s and its type is %d\n", t.text, t.tktype)
	}
	return
}
