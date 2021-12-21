package main

import (
	"fmt"
	"os"

	"github.com/carlosghabrous/gogym/pkg/gym/bluebook"
	"github.com/carlosghabrous/gogym/pkg/menu"
)

func init() {
	initBlueBook()
}

func initBlueBook() {
	e01 := &menu.Exercise{MD: menu.MetaData{Id: "Statistics", Description: "Chapter02, Ex01"}, Runner: bluebook.Exercise01}
	e02 := &menu.Exercise{MD: menu.MetaData{Id: "Quadratic", Description: "Chapter02, Ex02"}, Runner: bluebook.Exercise02}
	blue01 := &menu.Section{MD: menu.MetaData{Id: "Chapter02", Description: "Blue book's chapter 02"}}

	blue01.Attach(e01)
	blue01.Attach(e02)

	blueBook := &menu.Section{MD: menu.MetaData{Id: "Programming in Go book", Description: "Programming in Go book's exercises"}}
	blueBook.Attach(blue01)

	menu.Add("Programming in Go book", blueBook)
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
