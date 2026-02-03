package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// We will import internal packages here later
	// "github.com/Tikkaaa3/t-cli/internal/executor"
	// "github.com/Tikkaaa3/t-cli/internal/grader"
	// "github.com/Tikkaaa3/t-cli/internal/ui"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "t-cli [task_token]",
	Short: "Educational CLI Runner",
	Long:  `Executes your local code and verifies it against the learning platform requirements.`,

	// This enforces that exactly 1 argument (the Token/ID) is passed
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		taskToken := args[0]

		// 1. Start UI Spinner (Visuals)
		fmt.Println("Connecting to platform...")
		// ui.StartSpinner()

		// 2. Fetch Task Details (API)
		// Note: The token tells the backend WHO the user is and WHICH task this is.
		fmt.Printf("Fetching details for token: %s...\n", taskToken)
		// task, err := api.GetTask(taskToken)
		// if err != nil {
		// 	fmt.Errorf("%v", err)
		// }

		// 3. Execute Local Code (Executor)
		// fmt.Printf("⚡ Running command: %s\n", task.Steps.Command)
		// output, err := executor.Run(task.Steps.Command)

		// 4. Compare Results (Grader)
		// passed := grader.Check(output, task.Steps.ExpectedOutput)

		// 5. Submit Results (API)
		// api.SubmitResult(task.UserID, passed)

		// Temporary output to prove it works
		fmt.Println("\nWorkflow placeholder complete.")
		fmt.Printf("Received Token: %s\n", taskToken)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
