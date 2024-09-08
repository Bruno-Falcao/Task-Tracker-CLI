## What this project is capable of:
* Create a new task
* Update tasks
* Delete tasks
* Change tasks status
* list all tasks
* list tasks by status

## Requisites to run the project
 * Go SDK;

## Commands to run this project:
```
go run . add "your description to the task" 
go run . update integerID "your new description to the specified task"
go run . delete integerID
go run . mark-in-progress integerID
go run . mark-done integerID
go run . mark-todo integerID
go run . list
go run . list done
go run . list in-progress
go run . list todo
```

https://github.com/Bruno-Falcao/Task-Tracker-CLI/tree/master
