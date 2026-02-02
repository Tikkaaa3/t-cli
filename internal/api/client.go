package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CommandStep struct {
	Command        string `json:"command"`
	ExpectedOutput string `json:"expected_output"`
}

type Task struct {
	UserID string        `json:"user_id"`
	Steps  []CommandStep `json:"steps"`
}

func GetTask(taskToken string) (Task, error) {
	api := "weDontKnowYet"
	client := &http.Client{
		Timeout: 10 * time.Second, // This kills the request if it takes longer than 10s
	}

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return Task{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", taskToken))
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
	if err != nil {
		return Task{}, err
	}

	return task, nil
}
