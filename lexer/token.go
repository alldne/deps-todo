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

func (t Token) String() string {
	var typeStr string
	switch t.Type {
	case NAME:
		typeStr = "NAME"
	case HYPHEN:
		typeStr = "HYPHEN"
	case COLON:
		typeStr = "COLON"
	case COMMA:
		typeStr = "COMMA"
	}
	return fmt.Sprintf("Token(%s, %s)", typeStr, t.Text)
}
