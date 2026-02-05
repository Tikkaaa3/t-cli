package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define the "Fake Task" data
var mockTaskResponse = `
{
  "id": "task-101",
  "steps": [
    {
      "command": "echo Hello World",
      "expected_output": "Hello World"
    },
    {
      "command": "echo This is a test",
      "expected_output": "This is a test"
    }
  ]
}
`

type SubmissionRequest struct {
	TaskID string `json:"task_id"`
	Passed bool   `json:"passed"`
}

func main() {
	mux := http.NewServeMux()

	// --- GET /tasks/{id} ---
	// This simulates fetching the instructions for a specific task
	mux.HandleFunc("GET /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Extract the ID from the URL path
		taskID := r.PathValue("id")

		// Check headers
		token := r.Header.Get("Authorization")

		fmt.Printf("\n--- CLI Request: Fetch Task ---\n")
		fmt.Printf("Requested Task ID: %s\n", taskID)
		fmt.Printf("Auth Token:        %s\n", token)

		if token == "" {
			fmt.Println("Warning: No Authorization header received!")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: Missing token"))
			return
		}

		// Send back the instructions
		fmt.Println("Action: Sending mock task JSON...")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockTaskResponse))
	})

	// --- POST /submissions ---
	mux.HandleFunc("POST /submissions", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			fmt.Println("Rejected: No token provided")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Parse the JSON Body
		var req SubmissionRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Check the Status
		fmt.Printf("\n--- CLI Submission ---\n")
		fmt.Printf("Task ID:    %s\n", req.TaskID)
		fmt.Printf("Auth Token: %s\n", token)

		if req.Passed {
			fmt.Println("RESULT: PASSED! Updating user progress in DB...")
		} else {
			fmt.Println("RESULT: FAILED. Logging attempt...")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("Mock Backend running on http://localhost:8080")
	fmt.Println("   - GET  /tasks/{id}")
	fmt.Println("   - POST /submissions")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
