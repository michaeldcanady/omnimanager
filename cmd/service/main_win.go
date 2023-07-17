//go:build windows
// +build windows

package main

import (
	"bufio"
	"fmt"

	"github.com/Microsoft/go-winio"
)

const pipePath = `\\.\pipe\your_pipe_name`

func actualMain() int {
	// Create a named pipe listener
	listener, err := winio.ListenPipe(pipePath, nil)
	if err != nil {
		fmt.Println("Error creating named pipe listener:", err)
		return 1
	}
	defer listener.Close()

	// Accept incoming connections from the named pipe
	pipe, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection from named pipe:", err)
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
		return 1
	}
	return 0
}
