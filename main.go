package main

import (
	"fmt"
	"os"

	"github.com/carlosghabrous/gogym/pkg/gym/bluebook"
	"github.com/carlosghabrous/gogym/pkg/menu"
)

//TODO: Default top level item to avoid nil fields etc
//TOOD: Interface to add stuff?
func init() {
	e01 := &menu.Exercise{menu.SingleSection{Id: "1", Description: "Chapter01, Ex01"}, bluebook.Exercise01}
	blue01 := &menu.Section{menu.SingleSection{Id: "Chapter01", Description: "Blue book's chapter 01"}, nil}
	blue01.Attach(e01)

	menu.Add("blue book", blue01)
}

func main() {
	fmt.Println("Starting GoGym!")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("See ya!")
}

func run() error {
	if err := menu.Loop(); err != nil {
		return fmt.Errorf("error while executing main loop: %v", err)
	}

	return nil
}
