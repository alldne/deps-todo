package lexer

import "fmt"

type tokenType int

const (
	NAME tokenType = iota
	HYPHEN
	COLON
	COMMA
)

type Token struct {
	Type tokenType
	Text string
}
