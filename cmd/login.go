package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func runLogin(cmd *cobra.Command, args []string) {
	token := args[0]

	home, _ := os.UserHomeDir()
	configPath := filepath.Join(home, ".t-cli") // e.g., /Users/me/.t-cli

	err := os.WriteFile(configPath, []byte(token), 0o600)
	if err != nil {
		fmt.Println("Error saving login:", err)
		return
	}
	fmt.Println("Success! Login credentials saved.")
}

func getAPIToken() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(home, ".t-cli")

	data, err := os.ReadFile(configPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("you are not logged in. Run 't-cli login <token>' first")
	}

	return strings.TrimSpace(string(data)), nil
}

var loginCmd = &cobra.Command{
	Use:   "login [api_token]",
	Short: "Save your API token locally",
	Args:  cobra.ExactArgs(1),
	Run:   runLogin,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
