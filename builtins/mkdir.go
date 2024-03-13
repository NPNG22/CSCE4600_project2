package builtins

import (
	"fmt"
	"os"
)

// MakeNewDirectory creates a new directory with the given name
func MakeNewDirectory(args ...string) error {
	// Check if the number of arguments is exactly 1
	if len(args) != 1 {
		return fmt.Errorf("%w: Expected one argument (directory name)", ErrInvalidArgCount)
	}

	dirName := args[0]

	// Attempt to create the directory with the specified permissions
	if err := os.Mkdir(dirName, 0750); err != nil {
		// Check if the error is due to the directory already existing
		if os.IsExist(err) {
			fmt.Println("Directory already exists")
			return nil // Return nil to indicate that the operation was not an error
		}

		// Print an error message if the directory creation failed
		fmt.Println("Failed to create directory", err)
		return err
	}

	// Print a success message if the directory was created
	fmt.Println("Directory created")
	return nil
}
