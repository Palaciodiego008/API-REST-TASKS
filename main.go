package main

import (
	"time"
)

// Data
type task struct {
	ID            int       `json:"ID"`
	Title         string    `json:"Title"`
	Description   string    `json:"Description"`
	DateCompleted time.Time `json:"DateCompleted"`
	Completed     bool      `json:"Completed"`
}

type allTasks []task

// Persistence
var tasks = allTasks{
	{
		ID:            1,
		Title:         "Task One",
		Description:   "Cook the fish",
		DateCompleted: time.Now().AddDate(-2, -5, 12),
		Completed:     true,
	},
}

func main() {
	Routes()
}
