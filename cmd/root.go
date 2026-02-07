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
	Use:   "t-cli [task_id]",
	Short: "Educational CLI Runner",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lessonID := args[0]
		var task api.Task
		var err error

		// Auth Check
		authToken, err := getAPIToken()
		if err != nil {
			ui.PrintFail(err.Error())
			os.Exit(1)
		}

		// Fetch Task
		_ = spinner.New().
			Title(fmt.Sprintf("Fetching task for lesson '%s'...", lessonID)).
			Action(func() {
				task, err = api.GetTask(lessonID, authToken)
			}).
			Run()

		if err != nil {
			ui.PrintFail(fmt.Sprintf("Could not fetch task: %v", err))
			os.Exit(1)
		}

		// Execute Code
		results, err := executor.Run(task.Steps, func(currentCommand string) {
			fmt.Printf("  > Running: %s ...\n", currentCommand)
		})

		time.Sleep(500 * time.Millisecond)

		if err != nil {
			ui.PrintFail(fmt.Sprintf("Execution system failure: %v", err))
			os.Exit(1)
		}

		// Grade & Show Results
		passed := grader.Check(results, task.Steps)

		if passed {
			fmt.Println(ui.ResultBox.BorderForeground(ui.SuccessStyle.GetForeground()).Render(
				ui.SuccessStyle.Render("CONGRATULATIONS!\n") +
					"All steps executed correctly.",
			))

			// Submit Result
			fmt.Println()
			_ = spinner.New().
				Title("Saving progress...").
				Action(func() {
					err = api.SubmitResult(task.ID, authToken)
				}).
				Run()

			if err != nil {
				ui.PrintFail("Could not save to server (but you passed locally!)")
			} else {
				ui.PrintSuccess("Progress saved.")
			}

		} else {
			fmt.Println(ui.ResultBox.BorderForeground(ui.FailStyle.GetForeground()).Render(
				ui.FailStyle.Render("TEST FAILED\n") +
					"Your output did not match expectation.",
			))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
