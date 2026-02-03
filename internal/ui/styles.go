package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Define colors
var (
	colorGreen = lipgloss.Color("42")  // Success
	colorRed   = lipgloss.Color("160") // Failure
	colorGray  = lipgloss.Color("240") // Info
)

// Define standard text styles
var (
	// Success: Green Text + Bold
	SuccessStyle = lipgloss.NewStyle().
			Foreground(colorGreen).
			Bold(true).
			PaddingLeft(1)

	// Fail: Red Text + Bold
	FailStyle = lipgloss.NewStyle().
			Foreground(colorRed).
			Bold(true).
			PaddingLeft(1)

	// Box: A border around the final result
	ResultBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			MarginTop(1)
)

// Helper functions for printing nice messages
func PrintSuccess(msg string) {
	fmt.Println(SuccessStyle.Render("✅ " + msg))
}

func PrintFail(msg string) {
	fmt.Println(FailStyle.Render("❌ " + msg))
}

func PrintInfo(msg string) {
	fmt.Println(lipgloss.NewStyle().Foreground(colorGray).Render("ℹ️  " + msg))
}
