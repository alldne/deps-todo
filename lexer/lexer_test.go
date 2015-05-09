package lexer

import (
	"fmt"
	"time"

	"testing"
)

func checkTokenSequence(tokenChan chan Token, expectedTokenTypeSeq []tokenType) bool {
	i := 0
	for t := range tokenChan {
		if i >= len(expectedTokenTypeSeq) {
			fmt.Printf("token pushed more than expected\npushed: %s\n", t)
			return false
		}
		fmt.Printf("token type expected: %s, got: %s\n", expectedTokenTypeSeq[i], t.Type)
		if t.Type != expectedTokenTypeSeq[i] {
			return false
		}
		i = i + 1
	}
	if i != len(expectedTokenTypeSeq) {
		fmt.Printf("token pushed less than expected\nremain: %s\n", expectedTokenTypeSeq[i:])
		return false
	}
	return true
}

func TestRun(t *testing.T) {
	tokenChan := make(chan Token)
	src := "task 1: task 2, task 3 //comment \n- subtask 1\n- subtask 2\n\ntask 2\n\ntask 3\n- task 4 // comment\n"
	expectedTokenTypeSeq := []tokenType{NAME, COLON, NAME, COMMA, NAME, HYPHEN, NAME, HYPHEN, NAME, NAME, NAME, HYPHEN, NAME, EOF}

	go Run(tokenChan, &src)

	testResult := make(chan bool)

	go func() {
		testResult <- checkTokenSequence(tokenChan, expectedTokenTypeSeq)
	}()

	go func() {
		time.Sleep(time.Second * 5)
		testResult <- false
	}()

	if !(<-testResult) {
		t.Error()
	}
}
