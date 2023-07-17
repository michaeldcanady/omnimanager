//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os"
	"syscall"
)

const pipePath = "/path/to/your/named/pipe"

func actualMain() int {
	// Remove any existing file at the named pipe path
	os.Remove(pipePath)

	// Create the named pipe (FIFO)
	err := syscall.Mkfifo(pipePath, 0666)
	if err != nil {
		fmt.Println("Error creating named pipe:", err)
		return 1
	}
	defer os.Remove(pipePath)

	// Open the named pipe for reading
	pipe, err := os.OpenFile(pipePath, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		fmt.Println("Error opening named pipe:", err)
		return 1
	}
	defer pipe.Close()

	// Create a scanner to read data from the named pipe
	scanner := bufio.NewScanner(pipe)

	for scanner.Scan() {
		// Read the data from the named pipe
		data := scanner.Text()

		// Run the Cobra CLI application and pass the received data as arguments
		cmdArgs := []string{"myapp", data}
		fmt.Println("Running myapp with arguments:", cmdArgs)

		// You can execute the Cobra CLI application with the received data using os/exec package.
		// For this example, we will just print the command that would be executed.

		// Uncomment the following lines to execute the Cobra CLI app using os/exec
		/*
			cmd := exec.Command("myapp", data)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error running myapp:", err)
			}
		*/
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from named pipe:", err)
	}
	return 0
}
