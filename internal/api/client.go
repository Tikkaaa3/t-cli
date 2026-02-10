package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var BaseURL = "http://localhost:8080"

type CommandStep struct {
	Command        string `json:"command"`
	ExpectedOutput string `json:"expected_output"`
}

type Task struct {
	ID    string        `json:"task_id"`
	Steps []CommandStep `json:"steps"`
}

func GetTask(lessonID, authToken string) (Task, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	url := fmt.Sprintf("%s/lessons/%s/task", BaseURL, lessonID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Task{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := client.Do(req)
	if err != nil {
		return Task{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Task{}, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var task Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	return task, err
}

func SubmitResult(taskID, authToken string) error {
	client := &http.Client{Timeout: 10 * time.Second}

	url := fmt.Sprintf("%s/tasks/%s/complete", BaseURL, taskID)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("server rejected submission: status %d", resp.StatusCode)
	}
	return nil
}
