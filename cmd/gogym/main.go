package main

import (
	"fmt"
	"os"

	gym "github.com/carlosghabrous/gogym"
)

// func init() {
// 	initBlueBook()
// 	// initOrangeBook()
// }

// func initBlueBook() {
// 	// Chapter 02
// 	e01c02 := &menu.Exercise{MD: menu.MetaData{Id: "Statistics", Description: "Chapter02, Ex01"}, Runner: bluebook.Exercise01C02}
// 	e02c02 := &menu.Exercise{MD: menu.MetaData{Id: "Quadratic", Description: "Chapter02, Ex02"}, Runner: bluebook.Exercise02C02}
// 	blue02 := &menu.Section{MD: menu.MetaData{Id: "Chapter02", Description: "Ch 02: Booleans and numbers"}}

// 	blue02.Attach(e01c02)
// 	blue02.Attach(e02c02)

// 	// Chapter 03
// 	e01c03 := &menu.Exercise{MD: menu.MetaData{Id: "M3u2pls", Description: "Chapter03, Ex01"}, Runner: bluebook.Exercise01C03}
// 	blue03 := &menu.Section{MD: menu.MetaData{Id: "Chapter03", Description: "Ch 03: Strings"}}
// 	blue03.Attach(e01c03)

// 	blueBook := &menu.Section{MD: menu.MetaData{Id: "Programming in Go book", Description: "Programming in Go book's exercises"}}
// 	blueBook.Attach(blue02)
// 	blueBook.Attach(blue03)

// 	menu.Add("Programming in Go book", blueBook)
// }

// func initOrangeBook() {
// 	orangeBook := menu.addSection("The go's oranage book")
// 	chapter1 := orangeBook.addChapter(1, "chapter 1 title")
// 	chapter1.addExercise(1, "title", "description", functionToRunExercise1)
// 	chapter1.addExercise(2, "title", "description", functionToRunExercise2)

// }

func initBlueBook(g *gym.Gym) error {
	_, err := g.AddSection("Blue Book")
	if err != nil {
		return err
	}

	return nil
}

func initSections(g *gym.Gym) error {

	err := initBlueBook(g)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Starting GoGym!")
	gym := gym.NewGym()

	err := initSections(gym)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
		os.Exit(1)
	}

	gym.Print()

	// if err := run(gym); err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
	// 	os.Exit(1)
	// }

	fmt.Println("See ya!")
}

// func run() error {
// 	if err := menu.Loop(); err != nil {
// 		return fmt.Errorf("error while executing main loop: %v", err)
// 	}

// 	return nil
// }
