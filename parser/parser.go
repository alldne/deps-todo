package parser

import (
	"../lexer"
	"os"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	if os.Getenv("DEBUG") == "1" {
		log.Level = logrus.DebugLevel
	}
}

type Todo []TaskDecl

type TaskDecl struct {
	MainTask Task
	Subtasks []Subtask
}

type Task struct {
	TaskName string
	TaskDeps Dependencies
}

type Subtask struct {
	TaskName string
}

type Dependencies []string

func consume() {
	lookAhead = <-tokenChan
	log.Debugf("Got token %s, %s", lookAhead.Type, lookAhead.Text)
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
	var taskDecls []TaskDecl
	taskDecls = append(taskDecls, parseTaskDecl())
	for lookAhead.Type == lexer.NAME {
		taskDecls = append(taskDecls, parseTaskDecl())
	}
	log.Debugf("parsed todo: %s", taskDecls)
	return taskDecls
}

func parseTaskDecl() TaskDecl {
	mainTask := parseMainTask()
	var subs []Subtask
	if lookAhead.Type == lexer.HYPHEN {
		subs = parseSubtasks()
	}
	log.Debugf("parsed task decl: %s, %s", mainTask, subs)
	return TaskDecl{mainTask, subs}
}

func parseMainTask() Task {
	taskName := consumeTaskName()
	var deps Dependencies
	if lookAhead.Type == lexer.COLON {
		consume()
		deps = parseDependencies()
	}
	log.Debugf("parsed main task: %s, %s", taskName, deps)
	return Task{taskName, deps}
}

func parseSubtasks() []Subtask {
	var subs []Subtask
	for lookAhead.Type == lexer.HYPHEN {
		consume()
		subs = append(subs, Subtask{consumeTaskName()})
	}
	log.Debugf("parsed subtasks: %s", subs)
	return subs
}

func parseDependencies() Dependencies {
	var taskNames []string
	taskNames = append(taskNames, consumeTaskName())
	for lookAhead.Type == lexer.COMMA {
		consume()
		taskNames = append(taskNames, consumeTaskName())
	}
	log.Debugf("parsed dependencies: %s", taskNames)
	return taskNames
}

var tokenChan chan lexer.Token
var lookAhead lexer.Token

func Run(nodeChan chan Todo, lexerTokenChan chan lexer.Token) {
	tokenChan = lexerTokenChan
	consume()
	nodeChan <- parse()
}
