package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Tikkaaa3/t-cli/internal/api"
	"github.com/Tikkaaa3/t-cli/internal/executor"
	"github.com/Tikkaaa3/t-cli/internal/grader"
	"github.com/Tikkaaa3/t-cli/internal/ui"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "t-cli [task_token]",
	Short: "Educational CLI Runner",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskToken := args[0]
		var task api.Task
		var err error

		// --- Fetch Task ---
		_ = spinner.New().
			Title("Fetching task details...").
			Action(func() {
				time.Sleep(2 * time.Second)
				task, err = api.GetTask(taskToken)
			}).
			Run()

		if err != nil {
			ui.PrintFail(fmt.Sprintf("Could not fetch task: %v", err))
			os.Exit(1)
		}

		fmt.Println() // Add a blank line for spacing
		ui.PrintInfo(fmt.Sprintf("Task loaded. Executing %d step(s):", len(task.Steps)))

		// --- Execute Code ---
		// We pass a function that prints the command nicely
		results, err := executor.Run(task.Steps, func(currentCommand string) {
			// This prints: "  > Running: python main.py"
			fmt.Printf("  > Running: %s ...\n", currentCommand)
		})

		// One final small pause before results
		time.Sleep(500 * time.Millisecond)

		if err != nil {
			ui.PrintFail(fmt.Sprintf("Execution system failure: %v", err))
			os.Exit(1)
		}

		// --- Grade & Show Results ---
		passed := grader.Check(results, task.Steps)

		if passed {
			fmt.Println(ui.ResultBox.BorderForeground(ui.SuccessStyle.GetForeground()).Render(
				ui.SuccessStyle.Render("CONGRATULATIONS!\n") +
					"All steps executed correctly.",
			))
		} else {
			fmt.Println(ui.ResultBox.BorderForeground(ui.FailStyle.GetForeground()).Render(
				ui.FailStyle.Render("TEST FAILED\n") +
					"Your output did not match expectation.",
			))
		}

		// --- Submit Result ---
		fmt.Println()
		_ = spinner.New().
			Title("Saving progress...").
			Action(func() {
				time.Sleep(2 * time.Second)
				err = api.SubmitResult(taskToken, passed)
			}).
			Run()

		if err != nil {
			ui.PrintFail("Could not save to server (but you passed locally!)")
		} else {
			ui.PrintSuccess("Progress saved.")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
