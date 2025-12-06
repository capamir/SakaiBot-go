package utils

import (
	"fmt"
	"github.com/fatih/color"
)

// Info prints an informational message in blue with emoji
func Info(message string) {
	color.Set(color.FgCyan)
	fmt.Printf("ℹ️  %s\n", message)
	color.Unset()
}

// Success prints a success message in green with emoji
func Success(message string) {
	color.Set(color.FgGreen)
	fmt.Printf("✅ %s\n", message)
	color.Unset()
}

// Warning prints a warning message in yellow with emoji
func Warning(message string) {
	color.Set(color.FgYellow)
	fmt.Printf("⚠️  %s\n", message)
	color.Unset()
}

// Error prints an error message in red with emoji
func Error(message string) {
	color.Set(color.FgRed)
	fmt.Printf("❌ %s\n", message)
	color.Unset()
}

// Debug prints a debug message in magenta with emoji
func Debug(message string) {
	color.Set(color.FgMagenta)
	fmt.Printf("🔍 %s\n", message)
	color.Unset()
}
