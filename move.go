package main

import (
	"fmt"
	"os"
	"time"
)

var usersDesktopDirectory string = os.Getenv("HOME") + "/Desktop"
var date string = time.Now().Format("2006-01-02")

func Run() {
	var files []string = ListFiles(usersDesktopDirectory, false)
	var archiveDirectory string = usersDesktopDirectory + "/" + ArchiveDirectory
	var destinationDirectory string = usersDesktopDirectory + "/" + ArchiveDirectory + "/" + date

	// Create the directories
	CreateDirectories([]string{archiveDirectory, destinationDirectory})

	fmt.Println("Files to move: ", len(files))

	for _, file := range files {
		// Move the file
		err := os.Rename(usersDesktopDirectory+"/"+file, destinationDirectory+"/"+file)

		if err != nil {
			fmt.Println("Error moving file: ", err)
		} else {
			fmt.Println("Moved file: ", file)
		}
	}
}
