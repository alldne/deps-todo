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
}

type subtaskType int

const (
	ORDERED subtaskType = iota
	UNORDERED
)

type Subtask struct {
	Type subtaskType
	TaskName string
}

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
	if lookAhead.Type == lexer.HYPHEN || lookAhead.Type == lexer.ASTERISK {
		subs = parseSubtasks()
	}
	log.Debugf("parsed task decl: %s, %s", mainTask, subs)
	return TaskDecl{mainTask, subs}
}

func parseMainTask() Task {
	taskName := consumeTaskName()
	log.Debugf("parsed main task: %s", taskName)
	return Task{taskName}
}

func parseSubtasks() []Subtask {
	var subs []Subtask
	for lookAhead.Type == lexer.HYPHEN || lookAhead.Type == lexer.ASTERISK {
		var t subtaskType
		if lookAhead.Type == lexer.HYPHEN {
			t = ORDERED
		} else if lookAhead.Type == lexer.ASTERISK {
			t = UNORDERED
		}
		consume()
		subs = append(subs, Subtask{t, consumeTaskName()})
	}
	log.Debugf("parsed subtasks: %s", subs)
	return subs
}

var tokenChan chan lexer.Token
var lookAhead lexer.Token

func Run(nodeChan chan Todo, lexerTokenChan chan lexer.Token) {
	tokenChan = lexerTokenChan
	consume()
	nodeChan <- parse()
}
