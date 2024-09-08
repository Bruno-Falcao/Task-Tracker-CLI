package main

import (
	"time"
)

type Status string

const (
	TODO       = "todo"
	INPROGRESS = "in-progress"
	DONE       = "done"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (task *Task) NewTask() *Task {
	return &Task{
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   time.Now().Format("02-01-2006 15:04:05"),
	}
}
