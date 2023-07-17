package main

import (
	"fmt"
	"os"

	"github.com/michaeldcanady/omnimanageragent/internal/command"
)

func main() {
	os.Exit(actualMain())
}

func actualMain() int {
	err := command.Execute()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
