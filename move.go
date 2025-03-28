package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var usersDesktopDirectory string = os.Getenv("HOME") + "/Desktop"
var date string = time.Now().Format("2006-01-02")

func Run(stepThrough bool) {

	fmt.Println("Step: ", stepThrough)

	var files []string = ListFiles(usersDesktopDirectory, false)
	var archiveDirectory string = usersDesktopDirectory + "/" + ArchiveDirectory
	var destinationDirectory string = usersDesktopDirectory + "/" + ArchiveDirectory + "/" + date

	// Create the directories
	CreateDirectories([]string{archiveDirectory, destinationDirectory})

	fmt.Println("Files to move: ", len(files))

	for _, file := range files {

		if stepThrough {
			fmt.Println("[Y/n] to move file: ", file)
			var input string
			fmt.Scanln(&input)

			var formatedInput string = strings.ToLower(input)

			if formatedInput == "n" {
				fmt.Println("Skipping file: ", file)
				continue
			} else if formatedInput != "y" {
				fmt.Println("Invalid input, skipping file: ", file)
				continue
			}
		}

		// Move the file
		err := os.Rename(usersDesktopDirectory+"/"+file, destinationDirectory+"/"+file)

		if err != nil {
			fmt.Println("Error moving file: ", err)
		} else {
			fmt.Println("Moved file: ", file)
		}
	}
}
