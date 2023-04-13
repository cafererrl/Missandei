package main

import (
	"fmt"
	"os"
)

// Codes the types of terminal text outputs.
var (
	Prompt = Yellow
	Info   = Teal
	Error  = Red
)

// Define the terminal text outputs colors.
var (
	Red     = toColor("\033[1;31m%s\033[0m")
	Green   = toColor("\033[1;32m%s\033[0m")
	Yellow  = toColor("\033[1;33m%s\033[0m")
	Purple  = toColor("\033[1;34m%s\033[0m")
	Magenta = toColor("\033[1;35m%s\033[0m")
	Teal    = toColor("\033[1;36m%s\033[0m")
)

// Helper to style the Interpreter outputs.
func toColor(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// Helper to handle errors.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Flags path errors so that user may try to input the correct program.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
