package main

import (
	"log"

	"github.com/Microsoft/go-winio"
)

// Example client
func main() {
	pipePath := `\\.\pipe\your_pipe_name`
	f, err := winio.DialPipe(pipePath, nil)
	if err != nil {
		log.Fatalf("error opening pipe: %v", err)
	}
	defer f.Close()
	n, err := f.Write([]byte("sync"))
	if err != nil {
		log.Fatalf("write error: %v", err)
	}
	log.Println("wrote:", n)
}
