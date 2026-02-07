package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define the response format the CLI expects
type CLIResponse struct {
	ID    string `json:"id"`
	Steps []struct {
		Command        string `json:"command"`
		ExpectedOutput string `json:"expected_output"`
	} `json:"steps"`
}

func main() {
	mux := http.NewServeMux()

	// --- GET /lessons/{lesson_id}/task ---
	mux.HandleFunc("GET /lessons/{lesson_id}/task", func(w http.ResponseWriter, r *http.Request) {
		lessonID := r.PathValue("lesson_id")
		token := r.Header.Get("Authorization")

		fmt.Printf("\n--- CLI Request: Fetch Task ---\n")
		fmt.Printf("Requested Lesson ID: %s\n", lessonID)
		fmt.Printf("Auth Token:          %s\n", token)

		if token == "" {
			fmt.Println("Warning: No Authorization header received!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Mock Data: A task that requires echoing
		response := CLIResponse{
			ID: "task-mock-uuid-123", // The "Task ID" (different from Lesson ID)
			Steps: []struct {
				Command        string `json:"command"`
				ExpectedOutput string `json:"expected_output"`
			}{
				{
					Command:        "echo Hello World",
					ExpectedOutput: "Hello World",
				},
				{
					Command:        "echo Testing 123",
					ExpectedOutput: "Testing 123",
				},
			},
		}

		fmt.Println("Action: Sending mock task JSON...")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// --- POST /tasks/{task_id}/complete ---
	mux.HandleFunc("POST /tasks/{task_id}/complete", func(w http.ResponseWriter, r *http.Request) {
		taskID := r.PathValue("task_id")
		token := r.Header.Get("Authorization")

		fmt.Printf("\n--- CLI Submission: Task Complete ---\n")
		fmt.Printf("Task ID:    %s\n", taskID) // Should match "task-mock-uuid-123"
		fmt.Printf("Auth Token: %s\n", token)

		if token == "" {
			fmt.Println("Rejected: No token provided")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Println("RESULT: PASSED! Marking task as complete...")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	})

	fmt.Println("Mock Backend running on http://localhost:8080")
	fmt.Println("   - GET  /lessons/{id}/task")
	fmt.Println("   - POST /tasks/{id}/complete")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
