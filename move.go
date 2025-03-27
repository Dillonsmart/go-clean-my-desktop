package main

import (
	"fmt"
	"os"
	"time"
)

const ArchiveDirectory string = "Archive"

var usersDesktopDirectory string = os.Getenv("HOME") + "/Desktop"
var date string = time.Now().Format("2006-01-02")

func Run() {
	fmt.Println(date)
	fmt.Println("User's desktop directory: ", usersDesktopDirectory)

	var files []string = ListFiles(usersDesktopDirectory, false)
	var archiveDirectory string = usersDesktopDirectory + "/" + ArchiveDirectory
	var destinationDirectory string = usersDesktopDirectory + "/" + ArchiveDirectory + "/" + date

	// Create the archive directory
	archiveDirectoryError := os.Mkdir(archiveDirectory, 0755)

	if archiveDirectoryError != nil {
		fmt.Println("Error creating archive directory: ", archiveDirectoryError)
		return
	}

	// Create the destination directory
	dateDirectoryError := os.Mkdir(destinationDirectory, 0755)

	if dateDirectoryError != nil {
		fmt.Println("Error creating directory: ", dateDirectoryError)
		return
	}

	for _, file := range files {
		fmt.Println("File: ", file)

		// Move the file
		err := os.Rename(usersDesktopDirectory+"/"+file, destinationDirectory+"/"+file)

		if err != nil {
			fmt.Println("Error moving file: ", err)
		}
	}
}
