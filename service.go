package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const fileName = "tasks.json"

func AddNewTask(taskDescription string) {
	nextId, err := getNextID()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	var task Task
	task.Id = nextId
	task.Description = taskDescription
	task.CreatedAt = time.Now().Format("02-01-2006 15:04:05")
	task.Status = TODO

	jsonData, done := saveTask(task)
	if done {
		return
	}

	if writeFile(jsonData) {
		return
	}

	fmt.Println("Success")
}

func UpdateTask(taskID int, newDescription string) {
	data, err := ReadFile()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	for i, task := range data {
		if task.Id == taskID {
			data[i].Description = newDescription
			data[i].UpdatedAt = time.Now().Format("02-01-2006 15:04:05")
			break
		}
	}

	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	writeFile(jsonData)
}

func ReadFile() ([]Task, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("File not exist")
	}

	var data []Task
	json.Unmarshal((content), &data)
	return data, nil
}

func UpdateStatus(statusName Status, taskID int) {
	data, err := ReadFile()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	for i, task := range data {
		if task.Id == taskID {
			data[i].Status = statusName
		}
	}

	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	writeFile(jsonData)
}

func DeleteTask(taskID int) {
	data, err := ReadFile()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	for i, task := range data {
		if task.Id == taskID {
			data = append(data[:i], data[i+1:]...)
		}
	}

	jsonData, err := json.MarshalIndent(data, "", " ")
	writeFile(jsonData)
}

func ListTasksByStatus(status Status) ([]Task, error) {
	tasks, err := ReadFile()
	if err != nil {
		return []Task{}, err
	}

	var taskByStatus []Task
	for _, task := range tasks {
		if task.Status == status {
			taskByStatus = append(taskByStatus, task)
		}
	}
	return taskByStatus, nil
}

func writeFile(jsonData []byte) bool {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return true
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return false
	}
	return true
}

func saveTask(task Task) ([]byte, bool) {
	data, err := ReadFile()

	if err != nil {
		fmt.Println(err)
	}
	tasks := append(data, task)
	jsonData, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		fmt.Println("Error serializing file:", err)
		return nil, true
	}
	return jsonData, false
}

// id autoincrement
func getNextID() (int, error) {
	tasks, err := ReadFile()
	if err != nil {
		return 1, err
	}
	maxID := 0

	for _, task := range tasks {
		if task.Id > maxID {
			maxID = task.Id
		}
	}

	return maxID + 1, nil
}
