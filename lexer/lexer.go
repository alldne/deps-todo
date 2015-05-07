package lexer

import (
	"bytes"
	"fmt"
	"strings"
)

var index int = 0
var src []rune
var lookAhead string

func consume() {
	index += 1
	if !isEnd() {
		lookAhead = string(src[index])
	}
	return
}

func whitespace() {
	for lookAhead == " " {
		consume()
	}
}

func comment() {
	if lookAhead == "/" {
		consume()
		if lookAhead == "/" {
			for lookAhead != "\n" {
				consume()
			}
			return
		} else {
			panic(`"/" expected`)
		}
	}
}

func taskname() string {
	var buffer bytes.Buffer
	for lookAhead != "\n" && lookAhead != ":" && lookAhead != "," {
		if lookAhead == "/" {
			t := lookAhead
			consume()
			if lookAhead == "/" {
				for lookAhead != "\n" {
					consume()
				}
				break
			} else {
				buffer.WriteString(t)
			}
		} else {
			buffer.WriteString(lookAhead)
			consume()
		}
	}
	ret := strings.Trim(buffer.String(), " ")
	fmt.Println(ret)
	return ret
	//    return strings.Trim(buffer.String(), " ")
}

func isEnd() bool {
	return index >= len(src)
}

func Run(tokenChan chan Token, srcString *string) {
	src = []rune(*srcString)
	index = -1
	consume()
	for !isEnd() {
		switch lookAhead {
		case " ":
			whitespace()
			continue
		case "\n":
			consume()
			continue
		case "/":
			comment()
			continue
		case "-":
			consume()
			tokenChan <- Token{HYPHEN, "-"}
			continue
		default:
			name := taskname()
			tokenChan <- Token{NAME, name}
			continue
		}
	}

	tokenChan <- Token{EOF, ""}
	close(tokenChan)
}
