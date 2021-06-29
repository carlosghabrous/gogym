package main

import (
	"fmt"
	"os"

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
		return fmt.Errorf("Expected one input argument")
	}

	err := menu.Loop(os.Args[1])
}
