package executor

import (
	"os/exec"
	"strings"

	"github.com/Tikkaaa3/t-cli/internal/api"
)

// TODO: Next Step: The Executor (internal/executor/runner.go)
// Now we need to write the code that actually runs the Python etc. commands on the user's machine.
// 3. Execute Local Code (Executor)
// output, err := executor.Run(task.Steps)

func Run(steps []api.CommandStep) (output []string, err error) {
	var results []string

	for _, commandStep := range steps {
		commandStr := commandStep.Command
		parts := strings.Fields(commandStr)

		if len(parts) == 0 {
			continue
		}

		cmd := exec.Command(parts[0], parts[1:]...)
		outputBytes, execErr := cmd.CombinedOutput()

		outputStr := string(outputBytes)

		// Edge Case: System Error
		// In this case, outputBytes is empty but execErr has the info.
		if len(outputStr) == 0 && execErr != nil {
			outputStr = execErr.Error()
		}

		results = append(results, outputStr)
	}
	return results, nil
}
