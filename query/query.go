package query

import (
	"../parser"
)

type querier struct {
	root    parser.Todo
	taskMap map[string]task
}

func (q querier) GetTodo(goal string) []string {
	t, ok := q.taskMap[goal]
	if !ok {
		return []string{goal}
	}

	if len(t.deps) == 0 {
		if len(t.subtasks) == 0 {
			return []string{t.name}
		} else {
			return q.GetTodo(t.subtasks[0])
		}
	} else {
		acc := make([]string, 0)
		for _, taskName := range t.deps {
			acc = append(acc, q.GetTodo(taskName)...)
		}
		return acc
	}
}

type task struct {
	name     string
	deps     []string
	subtasks []string
}

func makeTaskMap(root parser.Todo) map[string]task {
	t := make(map[string]task)
	addTask(t, root)
	return t
}

func addTask(tmap map[string]task, node interface{}) {
	switch node := node.(type) {
	case parser.Todo:
		for _, decl := range node {
			addTask(tmap, decl)
		}
	case parser.TaskDecl:
		if tmap[node.MainTask.TaskName].name != "" {
			panic("task is declared more than once")
		}

		subtasks := make([]string, len(node.Subtasks))
		for i, e := range node.Subtasks {
			subtasks[i] = e.TaskName
		}

		var t = task{subtasks: subtasks}
		tmap[node.MainTask.TaskName] = t
		addTask(tmap, node.MainTask)
	case parser.Task:
		tname := string(node.TaskName)
		t := tmap[tname]
		t.name = node.TaskName
		if len(node.TaskDeps) > 0 {
			t.deps = make([]string, len(node.TaskDeps))
			for i, str := range node.TaskDeps {
				t.deps[i] = string(str)
			}
		}
		tmap[tname] = t
	}
}

func New(root parser.Todo) querier {
	t := makeTaskMap(root)
	return querier{root, t}
}
