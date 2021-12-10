package menu

import "testing"

// exercise01 := &menu.Exercise{Name: "ex01", Description: "ex01 desc", Runner: bluebook.Exercise01}
// menu.Add(exercise01)

// // Adding a section to the top menu is possible

// // Adding a section with children to the top menu is possible
// section02 := &menu.Section{Name: "section 02", Description: "section02 desc"}
// section02.Add(exercise01)
// menu.Add(section02)

func TestAddSingleSectionToMenu(t *testing.T) {
	sectionName := "section 01"
	ss01 := &SingleSection{Name: sectionName, Description: "section01 desc"}
	Add(sectionName, ss01)

	data, err := Get(sectionName)

	if err != nil {
		t.Fatalf("error getting singleSection %s, expected nil", sectionName)
	}
	if !Compare(ss01, data) {
		t.Fatalf("error; got %s, expected %s", data, ss01)
	}
}

func TestAddExerciseToMenu(t *testing.T) {
	exerciseName := "ex01"
	ex01 := &Exercise{SingleSection{Name: exerciseName, Description: "Ex description"},
		func(a, b int) error { return nil }}

	Add(exerciseName, ex01)
}

func TestAddSectionToMenu(t *testing.T) {
	sectionName := "section01"
	s01 := &Section{SingleSection{Name: sectionName, Description: "desc 01"},
		make(MenuChildren)}

	Add(sectionName, s01)
	data, err := Get(sectionName)
	if err != nil {
		t.Fatalf("error getting Section %s, expected nil", sectionName)
	}

	if !Compare(s01, data) {
		t.Fatalf("error; got %s, expected %s", data, s01)
	}
}
