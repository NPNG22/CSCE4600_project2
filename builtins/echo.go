package builtins

import (
	"fmt"
	"strings"
)

// Prints the provided arguments into the console
func Echo(args ...string) error {
	// Join the individual arguments into a single string with space
	message := strings.Join(args, " ")

	// Print the message into the console
	fmt.Println(message)

	return nil
}
