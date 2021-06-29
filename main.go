package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/carlosghabrous/gogym/pkg/menu"
)

func main() {
	fmt.Println("Starting GoGym!")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("See ya!")
}

func run() error {
	if len(os.Args[1:]) != 1 {
		return fmt.Errorf("%s %s", "Expected one input argument. Got ", strings.Join(os.Args[1:], ","))
	}

	if err := menu.Loop(os.Args[1]); err != nil {
		return fmt.Errorf("Error while executing main loop: %v\n", err)
	}

	return nil
}
