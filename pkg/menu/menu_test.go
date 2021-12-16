package menu

import (
	"fmt"
	"testing"
)

func cleanUp(t *testing.T) {
	t.Cleanup(func() {
		for k := range topMenu.Children {
			delete(topMenu.Children, k)
		}
	})
}

func TestEmptyMenu(t *testing.T) {
	if data, err := get(&options{name: "something"}); err == nil {
		t.Fatalf("expected error, got %s", data)
	}
}

func TestAddExerciseToMenu(t *testing.T) {
	exerciseName := "ex01"
	runner := func(args ...interface{}) error {
		fmt.Printf("Do something here with args %s", args)
		return nil
	}
	ex01 := &Exercise{MetaData{Id: exerciseName, Description: "Ex description"},
		runner}

	Add(exerciseName, ex01)

	data, err := get(&options{name: exerciseName})

	if err != nil {
		t.Fatalf("error getting MetaData %s, expected nil", exerciseName)
	}

	if !equal(ex01, data) {
		t.Fatalf("error; got %s, expected %s", data, ex01)
	}
	cleanUp(t)
}

func TestAddSectionToMenu(t *testing.T) {
	sectionName := "section01"
	s01 := &Section{MetaData{Id: sectionName, Description: "desc 01"},
		make(Offspring)}

	Add(sectionName, s01)
	data, err := get(&options{name: sectionName})
	if err != nil {
		t.Fatalf("error getting Section %s, expected nil", sectionName)
	}

	if !equal(s01, data) {
		t.Fatalf("error; got %s, expected %s", data, s01)
	}
	cleanUp(t)
}

func TestAddSectionWithExerciseToMenu(t *testing.T) {
	runner := func(args ...interface{}) error {
		a := args[0].(int)
		b := args[1].(int)

		fmt.Printf("result: %d + %d = %d", a, b, a+b)
		return nil
	}

	exercise01 := &Exercise{MetaData{Id: "1", Description: "First exercise"}, runner}
	sectionName := "section01"
	s01 := &Section{MD: MetaData{Id: sectionName, Description: "section 01"}}
	s01.Attach(exercise01)

	Add(sectionName, s01)

	s, err := get(&options{name: sectionName})
	if err != nil {
		t.Fatalf("section %s not found in top menu", sectionName)
	}

	e, err := get(&options{name: "1", from: s})
	if err != nil {
		t.Fatalf("exercise %s not in section %s", "1", sectionName)
	}

	if !equal(e, exercise01) {
		t.Fatalf("error; got %s, expected %s", e, exercise01)
	}

	cleanUp(t)
}

func TestAddNestedSectionsWithExerciseToMenu(t *testing.T) {
	runner := func(args ...interface{}) error {
		a := args[0].(int)
		b := args[1].(int)

		fmt.Printf("result: %d + %d = %d", a, b, a+b)
		return nil
	}

	exercise01 := &Exercise{MetaData{Id: "1", Description: "First exercise"}, runner}
	sectionName := "subsection 01"
	ss01 := &Section{MD: MetaData{Id: sectionName, Description: "subsection 01"}}
	ss01.Attach(exercise01)

	s01 := &Section{MD: MetaData{Id: "Main section", Description: "main section"}}
	s01.Attach(ss01)

	Add("Main section", s01)

	ms, err := get(&options{name: "Main section"})
	if err != nil {
		t.Fatalf("section %s not in top menu (%s)", ms, err)
	}

	cleanUp(t)

}
