package main

import (
	"fmt"
	"os"
)

const ArchiveDirectory string = "Archive"

func ListFiles(directory string, listHidden bool) []string {
	var moveableFiles []string = make([]string, 0)

	dir, err := os.Open(directory)
	if err != nil {
		fmt.Println("Error opening directory: ", err)
		return moveableFiles
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
		return moveableFiles
	}

	for _, file := range files {
		if file.Name()[0] == '.' && !listHidden || file.Name() == ArchiveDirectory {
			fmt.Println("Skipping file: ", file.Name())
			continue
		}

		fmt.Println(file.Name())
		moveableFiles = append(moveableFiles, file.Name())
	}

	return moveableFiles
}

func DirectoryExists(directory string) bool {
	info, err := os.Stat(directory)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println("Error checking directory:", err)
		return false
	}
	return info.IsDir()
}

func CreateDirectories(directories []string) {
	for _, directory := range directories {
		if !DirectoryExists(directory) {
			dateDirectoryError := os.Mkdir(directory, 0755)

			if dateDirectoryError != nil {
				fmt.Println("Error creating directory: ", dateDirectoryError)
				return
			}
		}
	}
}
