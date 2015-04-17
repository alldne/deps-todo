package lexer

type Token struct {
	Type tokenType
	Text string
}

type tokenType int

const (
	NAME tokenType = iota
	EQUAL
	HYPHEN
	COLON
)

func Run(tokenChan chan token, src string) {
	tokenChan <- token{EQUAL, "="}
	close(tokenChan)
}
