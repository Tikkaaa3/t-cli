package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define the specific "Fake Task" we want to test with
var mockTaskResponse = `
{
  "steps": [
    {
      "command": "echo Hello World",
      "expected_output": "Hello World"
    },
    {
      "command": "echo Hello World",
      "expected_output": "Hello World"
    },
    {
      "command": "echo Hello World",
      "expected_output": "Hello World"
    }
  ]
}
`

func handler(w http.ResponseWriter, r *http.Request) {
	// Log the request so you see what the CLI is sending
	fmt.Printf("\nReceived %s request from CLI\n", r.Method)

	// Check the Auth Token
	token := r.Header.Get("Authorization")
	fmt.Printf("Token: %s\n", token)

	if r.Method == "GET" {
		// --- SCENARIO 1: CLI is asking for the task ---
		fmt.Println("Sending mock task...")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockTaskResponse)) // Send the JSON
		return
	}

	if r.Method == "POST" {
		// --- SCENARIO 2: CLI is submitting results ---
		body, _ := io.ReadAll(r.Body)
		fmt.Printf("Submission Body: %s\n", string(body))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Mock Backend running on http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop...")

	// Listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
