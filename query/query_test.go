package query

import (
	"testing"

	"../parser"
)

func isSame(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAddTask(t *testing.T) {
	deps := parser.Dependencies{"d1", "d2", "d3"}
	ptask := parser.Task{"t", deps}
	taskDecl := parser.TaskDecl{ptask, []parser.Subtask{}}
	todo := parser.Todo([]parser.TaskDecl{taskDecl})

	var tmap map[string]task

	expected := []string{"d1", "d2", "d3"}

	tmap = make(map[string]task)
	addTask(tmap, ptask)
	if !isSame(tmap["t"].deps, expected) {
		t.Error("wrong result when Task given")
	}

	tmap = make(map[string]task)
	addTask(tmap, taskDecl)
	if !isSame(tmap["t"].deps, expected) {
		t.Error("wrong result when TaskDecl given")
	}

	tmap = make(map[string]task)
	addTask(tmap, todo)
	if !isSame(tmap["t"].deps, expected) {
		t.Error("wrong result when Todo given")
	}

	tmap = make(map[string]task)
	addTask(tmap, deps)
	if !isSame(tmap["t"].deps, []string{}) {
		t.Error("wrong result when wrong input given")
	}

	tmap = make(map[string]task)
	addTask(tmap, nil)
	if !isSame(tmap["t"].deps, []string{}) {
		t.Error("wrong result when wrong input given")
	}
}
