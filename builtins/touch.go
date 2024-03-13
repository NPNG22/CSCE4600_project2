package builtins

import (
	"fmt"
	"os"
	"time"
)

// Creates an empty file or updates its timestamps
func Touch(args ...string) error {
	// Check if the number of arguments is exactly 1
	if len(args) != 1 {
		return fmt.Errorf("%w: Expected one argument (file name)", ErrInvalidArgCount)
	}

	fileName := args[0]

	// Check if the file exists.
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// If the file does not exist, create an empty file
		if err := createEmptyFile(fileName); err != nil {
			return fmt.Errorf("Touch is unsuccessful: %s", err)
		}
		fmt.Println("File created")
	} else {
		// If the file exists, update its timestamps
		if err := updateTimestamps(fileName); err != nil {
			return fmt.Errorf("Touch is unsuccessful: %s", err)
		}
		fmt.Println("Timestamp updated")
	}

	return nil
}

// Creates an empty file with the given fileName
func createEmptyFile(fileName string) error {
	// Create the file.
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// Updates the access and modification timestamps of the file
func updateTimestamps(fileName string) error {
	// Get the current local time
	currentTime := time.Now().Local()

	// Update the timestamps of the file
	return os.Chtimes(fileName, currentTime, currentTime)
}
