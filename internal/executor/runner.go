package executor

import (
	"os/exec"
	"runtime"
	"time"

	"github.com/Tikkaaa3/t-cli/internal/api"
)

func Run(steps []api.CommandStep, onStep func(string)) (output []string, err error) {
	var results []string

	for _, commandStep := range steps {
		commandStr := commandStep.Command

		time.Sleep(500 * time.Millisecond)

		if onStep != nil {
			onStep(commandStr)
		}

		// Detect OS and wrap command in a shell
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			// On Windows, use "cmd /C"
			cmd = exec.Command("cmd", "/C", commandStr)
		} else {
			// On Mac/Linux, use "sh -c"
			cmd = exec.Command("sh", "-c", commandStr)
		}

		outputBytes, execErr := cmd.CombinedOutput()
		outputStr := string(outputBytes)

		// Edge Case: System Error
		if len(outputStr) == 0 && execErr != nil {
			outputStr = execErr.Error()
		}

		results = append(results, outputStr)
	}
	return results, nil
}
