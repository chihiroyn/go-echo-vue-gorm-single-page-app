package models

import "github.com/jinzhu/gorm"

// Task is a struct containing Task data
type Task struct {
//	ID   int    `json:"id"`
	gorm.Model
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks returns all tasks in database as TaskCollection
func GetTasks() (tc TaskCollection) {
	tasks := []Task{}
	db.Find(&tasks)

	tc = TaskCollection{
		Tasks: tasks,
	}

	return
}
