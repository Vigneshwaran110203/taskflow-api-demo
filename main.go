
package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	http.HandleFunc("/tasks/1", func(w http.ResponseWriter, r *http.Request) {
		task := Task{
			ID:    1,
			Title: "Prepare report",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	})

	http.ListenAndServe(":8080", nil)
}
