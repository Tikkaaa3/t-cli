package executor

import (
	"os/exec"
	"strings"

	"github.com/Tikkaaa3/t-cli/internal/api"
)

func Run(steps []api.CommandStep) (output []string, err error) {
	var results []string

	for _, commandStep := range steps {
		commandStr := commandStep.Command
		parts := strings.Fields(commandStr)

		if len(parts) == 0 {
			results = append(results, "")
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
