# deps-todo
tools for todo list, including simple domain specific language

# grammar
```
Todo = TaskDecl { "\n" TaskDecl }
TaskDecl = MainTask [ "\n" Subtasks ]
MainTask = TaskName
Subtasks = Subtask { "\n" SubTask }
SubTask = OrderedSubTask | UnorderedSubtask
OrderedSubTask = "-" TaskName
UnorderedSubtask = "*" TaskName
TaskName = <any character except newline>
```
