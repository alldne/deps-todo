package query

import (
	"../parser"
)

type querier struct {
	root parser.Todo
}

func New(root parser.Todo) querier {
	return querier{root}
}
