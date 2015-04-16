package main

type token struct {
	tktype tokenType
	text   string
}

type tokenType int

const (
	NAME tokenType = iota
	EQUAL
	HYPHEN
	COLON
)

func lexer(tokenChan chan token, src string) {
	tokenChan <- token{EQUAL, "="}
	close(tokenChan)
}
