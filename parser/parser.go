package parser

import "../lexer"

type Todo struct {
	taskDecls []taskDecl
}

type taskDecl struct {
	mainTask task
	subTasks []subtask
}

type task struct {
	taskName string
	taskDeps dependencies
}

type subtask struct {
	taskName string
}

type dependencies struct {
	taskNames []string
}

func Run(nodeChan chan Todo, tokenChan chan lexer.Token) {
	nodeChan <- Todo{[]taskDecl{}}
}
