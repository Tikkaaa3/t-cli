package api

type CommandStep struct {
	Command        string
	ExpectedOutput string
}

type Task struct {
	UserID string
	Steps  []CommandStep
}
