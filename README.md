# deps-todo
tools for todo list, including simple domain specific language

# grammar
```
Todo = TaskDecl { TaskDecl }
TaskDecl = MainTask [ Subtask ]
MainTask = TaskName [ ":" Dependencies ]
SubTask = "-" TaskName { "-" TaskName }
Dependencies = TaskName { "," TaskName }
TaskName = <any character except "-" or ":" or newline>
```
