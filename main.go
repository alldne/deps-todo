package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"./lexer"
	"./parser"
	"./print"
	"./query"
)

func main() {
	dat, err := ioutil.ReadFile("sample.dtodo")
	if err != nil {
		panic(err)
	}

	todo(string(dat))
}

func todo(src string) {
	tokenChan := make(chan lexer.Token)
	treeChan := make(chan parser.Todo)
	go lexer.Run(tokenChan, &src)
	go parser.Run(treeChan, tokenChan)
	root := <-treeChan
	fmt.Printf("Input:\n%s\n", print.Stringify(root))

	fmt.Println()

	q := query.New(root)

	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("To do .. (type something): ")
		line, _ := in.ReadString('\n')
		line = strings.Trim(line, "\n")
		if line == "" {
			continue
		}

		todos := q.GetTodo(line)
		if len(todos) == 1 {
			fmt.Printf("Do \"%s\"!\n", todos[0])
		} else {
			fmt.Println("Here are your todo list")
			for _, todo := range todos {
				fmt.Printf(" - %s\n", todo)
			}
		}
	}
	return
}
