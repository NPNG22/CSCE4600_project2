package builtins

import (
	"fmt"
	"os"
)

// Prints the current working directory
func PrintCurrentDirectory() error {
	// Retrieve the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		// If there's an error getting the current directory, print the error
		fmt.Printf("Error retrieving current directory: %s\n", err)
		return nil
	}

	// Print the current directory
	fmt.Printf("Current directory: %s\n", currentDir)
	return nil
}
