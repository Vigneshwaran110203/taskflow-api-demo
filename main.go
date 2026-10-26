package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	DueDate  string `json:"due_date"`
}

func main() {
	http.HandleFunc("/tasks/1", func(w http.ResponseWriter, r *http.Request) {
		task := Task{
			ID:       1,
			Title:    "Prepare report",
			Priority: "high",
			DueDate:  "2026-11-05",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	})

	http.ListenAndServe(":8080", nil)
}
