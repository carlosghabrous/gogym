package menu

import (
	"fmt"
	"reflect"
)

type MenuChildren map[string]Menuer
type SingleSection struct {
	Name        string
	Description string
}

type Section struct {
	SingleSection
	Children MenuChildren
}

type Exercise struct {
	SingleSection
	Runner func(args ...interface{}) error
}

type Menuer interface {
	String() string
	Add(menuer *Menuer)
}

func (ss *SingleSection) String() string {
	return fmt.Sprintf("Name: %s - Description: %s", ss.Name, ss.Description)
}

func (ss *SingleSection) Add(menuer *Menuer) {

}

func (s *Section) String() string {
	return fmt.Sprintf("Name: %s - Description: %s - Children: %s", s.Name, s.Description, s.Children)
}

func (s *Section) Add(menuer *Menuer) {

}

func (e *Exercise) String() string {
	return fmt.Sprintf("Name: %s - Description: %s", e.Name, e.Description)
}

func (e *Exercise) Add(menuer *Menuer) {

}

var topMenu MenuChildren

func Add(name string, item Menuer) {

	if IsMenuEmpty() {
		topMenu = make(MenuChildren)
	}

	if _, ok := topMenu[name]; ok {
		panic("item already in menu")
	}

	topMenu[name] = item
}

func Get(name string) (Menuer, error) {
	var menuer Menuer

	menuer, ok := topMenu[name]
	if !ok {
		return menuer, fmt.Errorf("top menu doesn't contain item %s", name)
	}

	return menuer, nil
}

func IsMenuEmpty() bool {
	return topMenu == nil
}

func Display() {

}

func Equal(a, b Menuer) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		fmt.Println("different types")
		return false
	}

	switch a.(type) {
	case *Exercise:
		return areExercisesEqual(a.(*Exercise), b.(*Exercise))

	case *Section:
		return areSectionsEqual(a.(*Section), b.(*Section))

	case *SingleSection:
		return areSingleSectionsEqual(a.(*SingleSection), b.(*SingleSection))
	}
	return false
}

func areExercisesEqual(a, b *Exercise) bool {
	return a.Name == b.Name &&
		a.Description == b.Description
	//TODO: Compare functions?
	// &&
	// a.Runner == b.Runner
}

func areSingleSectionsEqual(a, b *SingleSection) bool {
	return a.Name == b.Name &&
		a.Description == b.Description
}
func areSectionsEqual(a, b *Section) bool {
	return a.Name == b.Name &&
		a.Description == b.Description
	//TODO: Compare MenuChildren
	// &&
	// a.Children == b.Children
}

// Implements the main loop for gogym
func Loop() error {
	fmt.Println()
	return nil
}
