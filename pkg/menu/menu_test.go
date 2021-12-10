package menu

import (
	"fmt"
	"testing"
)

func TestAddSingleSectionToMenu(t *testing.T) {
	singleSectionName := "section 01"
	ss01 := &SingleSection{Id: singleSectionName, Description: "section01 desc"}
	Add(singleSectionName, ss01)

	data, err := Get(singleSectionName)

	if err != nil {
		t.Fatalf("error getting singleSection %s, expected nil", singleSectionName)
	}

	if !Equal(ss01, data) {
		t.Fatalf("error; got %s, expected %s", data, ss01)
	}
}

func TestAddExerciseToMenu(t *testing.T) {
	exerciseName := "ex01"
	runner := func(args ...interface{}) error {
		fmt.Printf("Do something here with args %s", args)
		return nil
	}
	ex01 := &Exercise{SingleSection{Id: exerciseName, Description: "Ex description"},
		runner}

	Add(exerciseName, ex01)
}

func TestAddSectionToMenu(t *testing.T) {
	sectionName := "section01"
	s01 := &Section{SingleSection{Id: sectionName, Description: "desc 01"},
		make(MenuChildren)}

	Add(sectionName, s01)
	data, err := Get(sectionName)
	if err != nil {
		t.Fatalf("error getting Section %s, expected nil", sectionName)
	}

	if !Equal(s01, data) {
		t.Fatalf("error; got %s, expected %s", data, s01)
	}
}

// func TestAddSectionWithExerciseToMenu(t *testing.T) {

// }

// func TestAddSectionSubWithExerciseToMenu(t *testing.T) {

// }

// // // Adding a section with children to the top menu is possible
// // section02 := &menu.Section{Name: "section 02", Description: "section02 desc"}
// // section02.Add(exercise01)
// // menu.Add(section02)
