package main

import (
	"fmt"
	"os"

	"github.com/carlosghabrous/gogym/pkg/gym/bluebook"
	"github.com/carlosghabrous/gogym/pkg/menu"
)

func init() {
	e01 := &menu.Exercise{menu.SingleSection{Id: "1", Description: "Chapter01, Ex01"}, bluebook.Exercise01}
	e02 := &menu.Exercise{menu.SingleSection{Id: "2", Description: "Chapter01, Ex02"}, bluebook.Exercise02}
	blue01 := &menu.Section{menu.SingleSection{Id: "Chapter01", Description: "Blue book's chapter 01"}, nil}

	blue01.Attach(e01)
	blue01.Attach(e02)

	blueBook := &menu.Section{menu.SingleSection{Id: "Go's blue book", Description: "Blue book's exercises"}, nil}
	blueBook.Attach(blue01)

	menu.Add("Go's blue book", blueBook)
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
