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
	if Compare(item, &Section{}) {
		topMenu = make(MenuChildren)
	}

	if _, ok := topMenu[name]; ok {
		panic("item already in menu")
	}

	topMenu[name] = item
}

func Get(name string) (Menuer, error) {
	return &Section{}, nil
}

func Compare(a, b Menuer) bool {
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		return false
	}

	switch a.(type) {
	case *Exercise:
		return compareExercises(a.(*Exercise), b.(*Exercise))

	case *Section:
		return compareSections(a.(*Section), b.(*Section))

	case *SingleSection:
		return compareSingleSections(a.(*SingleSection), b.(*SingleSection))
	}
	return false
}

func compareExercises(a, b *Exercise) bool {
	return true
}

func compareSingleSections(a, b *SingleSection) bool {
	return true
}
func compareSections(a, b *Section) bool {
	return true
}

// Implements the main loop for gogym
func Loop() error {
	fmt.Println()
	return nil
}
