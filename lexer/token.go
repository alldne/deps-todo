package lexer

import "fmt"

type tokenType int

func (t tokenType) String() string {
	var typeStr string
	switch t {
	case NAME:
		typeStr = "NAME"
	case HYPHEN:
		typeStr = "HYPHEN"
	case COLON:
		typeStr = "COLON"
	case COMMA:
		typeStr = "COMMA"
	case EOF:
		typeStr = "EOF"
	}
	return typeStr
}

const (
	NAME tokenType = iota
	HYPHEN
	COLON
	COMMA
	EOF
)

type Token struct {
	Type tokenType
	Text string
}

func (t Token) String() string {
	typeStr := t.Type.String()
	return fmt.Sprintf("Token(%s, %s)", typeStr, t.Text)
}
