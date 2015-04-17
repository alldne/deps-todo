package parser

import "../lexer"

type Todo []taskDecl

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

type dependencies []string

func consume() {
	lookAhead = <-tokenChan
	return
}

func consumeTaskName() string {
	if lookAhead.Type != lexer.NAME {
		panic("name expected")
	}
	taskName := lookAhead.Text
	consume()
	return taskName
}

func parse() Todo {
	todo := parseTodo()
	if lookAhead.Type != lexer.EOF {
		panic("EOF expected")
	}
	return todo
}

func parseTodo() Todo {
	var taskDecls []taskDecl
	taskDecls = append(taskDecls, parseTaskDecl())
	for lookAhead.Type == lexer.NAME {
		taskDecls = append(taskDecls, parseTaskDecl())
	}
	return taskDecls
}

func parseTaskDecl() taskDecl {
	mainTask := parseMainTask()
	var subs []subtask
	if lookAhead.Type == lexer.HYPHEN {
		subs = parseSubtasks()
	}
	return taskDecl{mainTask, subs}
}

func parseMainTask() task {
	taskName := consumeTaskName()
	var deps dependencies
	if lookAhead.Type == lexer.COLON {
		consume()
		deps = parseDependencies()
	}
	return task{taskName, deps}
}

func parseSubtasks() []subtask {
	var subs []subtask
	for lookAhead.Type == lexer.HYPHEN {
		consume()
		subs = append(subs, subtask{consumeTaskName()})
	}
	return subs
}

func parseDependencies() dependencies {
	var taskNames []string
	taskNames = append(taskNames, consumeTaskName())
	for lookAhead.Type == lexer.COMMA {
		consume()
		taskNames = append(taskNames, consumeTaskName())
	}
	return taskNames
}

var tokenChan chan lexer.Token
var lookAhead lexer.Token

func Run(nodeChan chan Todo, lexerTokenChan chan lexer.Token) {
	tokenChan = lexerTokenChan
	nodeChan <- parse()
}
