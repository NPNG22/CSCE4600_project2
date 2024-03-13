package builtins

import (
	"fmt"
	"os"
	"path/filepath"
)

// Listing the contents of the specified directories
func ListAllFiles(args ...string) error {
	// If no directories are provided, list the contents of the current directory
	if len(args) == 0 {
		args = append(args, ".")
	}

	// Iterate through the provided directories and list their contents
	for _, dir := range args {
		fmt.Printf("Listing contents of directory: %s\n", dir)
		err := listDirectoryContents(dir)
		if err != nil {
			fmt.Printf("Error listing directory: %s\n", err)
			return err
		}
	}

	return nil
}

// Lists the contents of the specified directory
func listDirectoryContents(dir string) error {
	// Open the directory.
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// Iterate through the entries and print their names
	for _, entry := range dirEntries {
		if entry.Type()&os.ModeSymlink != 0 {
			linkDest, linkErr := os.Readlink(filepath.Join(dir, entry.Name()))
			if linkErr == nil {
				fmt.Printf("%s -> %s\n", entry.Name(), linkDest)
			} else {
				fmt.Printf("%s (broken symlink)\n", entry.Name())
			}
		} else {
			fmt.Println(entry.Name())
		}
	}

	return nil
}
