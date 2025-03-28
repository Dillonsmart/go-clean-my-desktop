# go-clean-my-desktop

A simple Go application to organize your desktop by moving files into an archive folder based on the current date.

## Features

- Automatically organizes files on your desktop into a dated archive folder.
- Skips hidden files unless explicitly specified.
- Creates necessary directories if they do not exist.
- **Step-through mode**: Allows you to confirm each file before moving it.


## Installation

1. Ensure you have Go installed on your system (version 1.24.1 or later).
2. Clone this repository:
   ```sh
   git clone https://github.com/Dillonsmart/go-clean-my-desktop.git
3. Navigate to the project directory:
`cd go-clean-my-desktop`
4. Build the application:
`go build -o clean-desktop`

## Usage
Run the application with the following command:

`./clean-desktop`

This will:

- Scan your desktop for files.
- Create an Archive folder on your desktop if it doesn't already exist.
- Move all non-hidden files into a subfolder within - Archive, named after the current date (e.g., Archive/2023-10-01).

### Step-through Mode
To enable step-through mode, run the application with the `-step` flag:
```./clean-desktop -step```

In this mode, the application will prompt you to confirm each file before moving it:

- Enter Y to move the file.
- Enter N to skip the file.
- Any other input will be treated as invalid, and the file will be skipped.