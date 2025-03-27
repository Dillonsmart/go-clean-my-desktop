package main

import (
	"fmt"
	"os"
)

func ListFiles(directory string, listHidden bool) []string {
	var moveableFiles []string = make([]string, 0)

	// Open the directory
	dir, err := os.Open(directory)
	if err != nil {
		fmt.Println("Error opening directory: ", err)
		return moveableFiles
	}
	defer dir.Close()

	// Read the files in the directory
	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
		return moveableFiles
	}

	// Print the files
	for _, file := range files {

		// check if the file is hidden
		if file.Name()[0] == '.' && !listHidden && file.Name() != ArchiveDirectory {
			continue
		}

		moveableFiles = append(moveableFiles, file.Name())
	}

	return moveableFiles
}
