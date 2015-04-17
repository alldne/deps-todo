package lexer

func Run(tokenChan chan token, src string) {
	tokenChan <- token{EQUAL, "="}
	close(tokenChan)
}
