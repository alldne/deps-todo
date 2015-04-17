package print

import (
	"fmt"
	"strings"

	"../parser"
)

func Stringify(t parser.Todo) string {
	return todo(t)
}

func todo(t parser.Todo) string {
	var strs []string
	for _, td := range t {
		strs = append(strs, taskDecl(td))
	}
	return strings.Join(strs, "\n\n")
}

func taskDecl(td parser.TaskDecl) string {
	mt := mainTask(td.MainTask)
	if len(td.Subtasks) > 0 {
		st := subtasks(td.Subtasks)
		return fmt.Sprintf("%s\n%s", mt, st)
	}
	return fmt.Sprintf("%s", mt)
}

func mainTask(mt parser.Task) string {
	if len(mt.TaskDeps) > 0 {
		return fmt.Sprintf("%s: %s", mt.TaskName, dependencies(mt.TaskDeps))
	}
	return mt.TaskName
}

func subtasks(st []parser.Subtask) string {
	var strs []string
	for _, sub := range st {
		strs = append(strs, fmt.Sprintf("- %s", sub.TaskName))
	}
	return strings.Join(strs, "\n")
}

func dependencies(d parser.Dependencies) string {
	return strings.Join(d, ", ")
}
