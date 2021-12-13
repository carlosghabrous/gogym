package menu

import (
	"fmt"
	"reflect"
)

type MenuChildren map[string]Menuer
type SingleSection struct {
	Id          string
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
	Name() string
	Attach(menuer Menuer) error
}

type Options struct {
	Name string
	From Menuer
}

func (ss *SingleSection) String() string {
	return fmt.Sprintf("Name: %s - Description: %s", ss.Id, ss.Description)
}

func (ss *SingleSection) Name() string {
	return ss.Id
}

func (ss *SingleSection) Attach(menuer Menuer) error {
	panic("single sections are not allowed to have children")
}

func (s *Section) String() string {
	return fmt.Sprintf("Name: %s - Description: %s - Children: %s", s.Id, s.Description, s.Children)
}

func (s *Section) Name() string {
	return s.Id
}

func (s *Section) Attach(menuer Menuer) error {
	if s.Children == nil {
		s.Children = make(MenuChildren)
	}

	name := menuer.Name()

	if _, ok := s.Children[name]; ok {
		return fmt.Errorf("item %s already exists", name)
	}

	s.Children[name] = menuer
	return nil
}

func (e *Exercise) String() string {
	return fmt.Sprintf("Name: %s - Description: %s", e.Id, e.Description)
}

func (e *Exercise) Name() string {
	return e.Id
}

func (e *Exercise) Attach(menuer Menuer) error {
	panic("exercises are not allowed to have children")
}

var topMenu Section

func Add(name string, item Menuer) {

	if IsMenuEmpty() {
		topMenu.Id = "GoGym"
		topMenu.Description = "Exercising in Go"
		topMenu.Children = make(MenuChildren)
	}

	if _, ok := topMenu.Children[name]; ok {
		panic("item already in menu")
	}

	topMenu.Children[name] = item
}

func Get(options *Options) (Menuer, error) {
	var returnMenuer, fromMenuer Menuer

	if options.From == nil {
		fromMenuer = &topMenu

	} else {
		fromMenuer = options.From
	}

	fromSection, ok := fromMenuer.(*Section)
	if !ok {
		fmt.Errorf("only Section(s) contain children elements")
	}

	returnMenuer, ok = fromSection.Children[options.Name]
	if !ok {
		return returnMenuer, fmt.Errorf("top menu doesn't contain item %s", options.Name)
	}

	return returnMenuer, nil
}

func IsMenuEmpty() bool {
	return topMenu.Children == nil
}

func Display() {
	fmt.Printf("topMenu: %v\n", topMenu)
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
	return a.Id == b.Id &&
		a.Description == b.Description
	//TODO: Compare functions?
	// &&
	// a.Runner == b.Runner
}

func areSingleSectionsEqual(a, b *SingleSection) bool {
	return a.Id == b.Id &&
		a.Description == b.Description
}
func areSectionsEqual(a, b *Section) bool {
	return a.Id == b.Id &&
		a.Description == b.Description
	//TODO: Compare MenuChildren
	// &&
	// a.Children == b.Children
}

// Implements the main loop for gogym
func Loop() error {
	Display()
	return nil
}
