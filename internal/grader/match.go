package grader

import (
	"strings"

	"github.com/Tikkaaa3/t-cli/internal/api"
)

func Check(output []string, steps []api.CommandStep) bool {
	if len(output) != len(steps) {
		return false
	}
	for i, step := range steps {
		actual := strings.TrimSpace(output[i])
		expected := strings.TrimSpace(step.ExpectedOutput)

		// Compare
		if actual != expected {
			return false
		}
	}
	return true
}
