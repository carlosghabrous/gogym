package menu

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang-collections/collections/stack"
)

const (
	exitOptionName        = "Exit"
	exitOptionValue       = 0
	exitOptionDescription = "Exit GoGym"
	backOptionName        = "Back"
	backOptionDescription = "Back one level"
)

// offspring contains the children for a menu item
type offspring map[string]nameDescriptioner

// numberedOffspring associates assings a number to each children of a section, in the order of insertion
type numberedOffspring map[int]string

// MetaData contains the minimum data for a Section or Exercise
type MetaData struct {
	Id          string // Item identifier
	Description string // Item description
}

// Section contains metadata and children bound to it
type Section struct {
	MD               MetaData
	Children         offspring
	NumberedChildren numberedOffspring
	NChildren        int
}

// Exercise contains metadata and a function that implements the exercise's code
type Exercise struct {
	MD     MetaData
	Runner func(args ...interface{}) error
}

type nameDescriptioner interface {
	Name() string
	Desc() string
}

type attacher interface {
	Attach(menuer MenuItemer) error
}

// TODO: Change its name, it's horrible, frankly
// MenuItemer is the interface implemented by the types MetaData, Section y Exercise
type MenuItemer interface {
	nameDescriptioner
	attacher
}

type options struct {
	name string
	from nameDescriptioner
}

type menuOptionTuple struct {
	name        string
	description string
	value       int
}

// String returns the string representation of a MetaData
func (ss *MetaData) String() string {
	return fmt.Sprintf("<Name: %s> <Description: %s>", ss.Id, ss.Description)
}

// Name returns the MetaData's name
func (ss *MetaData) Name() string {
	return ss.Id
}

func (ss *MetaData) Desc() string {
	return ss.Description
}

// String returns the string representation of a Section
func (s *Section) String() string {
	return fmt.Sprintf("<Name: %s> <Description: %s> <Children: %s>", s.MD.Id, s.MD.Description, s.Children)
}

// Name returns the Section's name
func (s *Section) Name() string {
	return s.MD.Id
}

func (s *Section) Desc() string {
	return s.MD.Description
}

// Attach binds a MetaData, Section or Exercise to a Section
//TODO: refactor this and Add (Add should call Attach)
func (s *Section) Attach(item nameDescriptioner) error {
	if s.Children == nil {
		s.Children = make(offspring)
		s.NumberedChildren = make(numberedOffspring)

		addExtraOption(s, &menuOptionTuple{name: exitOptionName,
			description: exitOptionDescription,
			value:       exitOptionValue})

		addExtraOption(s, &menuOptionTuple{name: backOptionName,
			description: backOptionDescription,
			value:       1})
	}

	name := item.Name()

	if _, ok := s.Children[name]; ok {
		return fmt.Errorf("item %s already exists", name)
	}

	s.Children[name] = item
	s.NChildren++

	if _, ok := s.NumberedChildren[s.NChildren]; ok {
		if s.NumberedChildren[s.NChildren] == backOptionName {
			s.NumberedChildren[s.NChildren+1] = backOptionName

		} else {
			return fmt.Errorf("item %s with option number %d already exists", s.NumberedChildren[s.NChildren], s.NChildren)
		}
	}

	s.NumberedChildren[s.NChildren] = name

	return nil
}

// String returns the string representation of as Exercise
func (e *Exercise) String() string {
	return fmt.Sprintf("<Name: %s> <Description: %s>", e.MD.Id, e.MD.Description)
}

// Name returns the Exercise's name
func (e *Exercise) Name() string {
	return e.MD.Id
}

func (e *Exercise) Desc() string {
	return e.MD.Description
}

func (o *offspring) String() string {
	result := make([]string, 0)
	for _, v := range *o {
		result = append(result, fmt.Sprintf("<Name: %s> <Description: %s>", v.Name(), v.Desc()))
	}

	return strings.Join(result, "\n")
}

// topMenu is the variable storing the whole menu
var topMenu Section

// Add adds an item of name 'name' to another item of type MenuItemer
func Add(name string, item nameDescriptioner) {

	if isSectionEmpty(&topMenu) {
		topMenu.MD.Id = "GoGym"
		topMenu.MD.Description = "Exercising in Go"
		topMenu.Children = make(offspring)
		topMenu.NumberedChildren = make(numberedOffspring)

		addExtraOption(&topMenu, &menuOptionTuple{name: exitOptionName,
			description: exitOptionDescription,
			value:       exitOptionValue})
	}

	if _, ok := topMenu.Children[name]; ok {
		panic("item already in menu")
	}

	topMenu.Children[name] = item
	topMenu.NChildren++
	topMenu.NumberedChildren[topMenu.NChildren] = name
}

// Implements the main loop for gogym
func Loop() error {
	var userOption int
	var menuStack *stack.Stack = stack.New()
	var theChosenOne interface{}
	var visitingSection *Section = &topMenu

	for {
		display(visitingSection)

		userOption = getUserOption(visitingSection)

		if userChoseExit(userOption) {
			break
		}

		if userChoseBack(userOption, visitingSection) {
			theChosenOne = menuStack.Pop()

		} else {
			optionName := visitingSection.NumberedChildren[userOption]
			theChosenOne = visitingSection.Children[optionName]
			menuStack.Push(visitingSection)
		}

		switch theChosenOne := theChosenOne.(type) {

		case *Exercise:
			theChosenOne.Runner()

		case *Section:
			visitingSection = theChosenOne

		default:
			panic("wadafak!")
		}
	}

	return nil
}

// options is used as the argument type for the get function,
// since the from argument can be optional
func addExtraOption(s *Section, options *menuOptionTuple) {
	s.Children[options.name] = &Section{MD: MetaData{Id: options.name, Description: options.description}}
	s.NumberedChildren[options.value] = options.name
}

// get is a helper that returns an element of type MenuItemer that has already been added to the menu, or error if the item is not found
func get(options *options) (nameDescriptioner, error) {
	var returnMenuer, fromMenuer nameDescriptioner

	if options.from == nil {
		fromMenuer = &topMenu

	} else {
		fromMenuer = options.from
	}

	fromSection, ok := fromMenuer.(*Section)
	if !ok {
		return nil, fmt.Errorf("%s", "only Section(s) contain children elements")
	}

	returnMenuer, ok = fromSection.Children[options.name]
	if !ok {
		return returnMenuer, fmt.Errorf("top menu doesn't contain item %s", options.name)
	}

	return returnMenuer, nil
}

// isSectionEmpty returns true if the section has not been initialized yet
func isSectionEmpty(s *Section) bool {
	return s.MD == MetaData{} && s.Children == nil
}

// equal returns true if two variables of the MenuItemer interface are equal, false otherwise
func equal(a, b nameDescriptioner) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		fmt.Println("different types")
		return false
	}

	switch a.(type) {
	case *Exercise:
		return areExercisesEqual(a.(*Exercise), b.(*Exercise))

	case *Section:
		return areSectionsEqual(a.(*Section), b.(*Section))

	case *MetaData:
		return areSingleSectionsEqual(a.(*MetaData), b.(*MetaData))
	}
	return false
}

// areExercisesEqual compares to variables of type Exercise
func areExercisesEqual(a, b *Exercise) bool {
	return a.MD.Id == b.MD.Id &&
		a.MD.Description == b.MD.Description
	//TODO: Compare functions?
	// &&
	// a.Runner == b.Runner
}

// areSingleSectionsEqual compares to variables of type SingleSections
func areSingleSectionsEqual(a, b *MetaData) bool {
	return a.Id == b.Id &&
		a.Description == b.Description
}

// areSectionsEqual compares to variables of type Section
func areSectionsEqual(a, b *Section) bool {
	return a.MD.Id == b.MD.Id &&
		a.MD.Description == b.MD.Description
	//TODO: Compare offspring
	// &&
	// a.Children == b.Children
}

// display shows the menu to the user
func display(section *Section) {
	var name, description, formatString string
	formatString = "-> %-10d%-20s%-20s\n"

	for i := 1; i < len(section.NumberedChildren); i++ {
		name = section.NumberedChildren[i]
		description = section.Children[name].Desc()
		fmt.Printf(formatString, i, name, description)
	}

	name = section.NumberedChildren[exitOptionValue]
	description = section.Children[name].Desc()
	fmt.Printf(formatString, exitOptionValue, name, description)
}

func getUserOption(s *Section) int {
	var userOption int

	for {

		fmt.Println("Select an option from the menu")

		_, err := fmt.Scanf("%d", &userOption)
		if err != nil {
			fmt.Println("Something went wrong! Try again!")
			continue
		}

		min, max := getValidRange(s)
		if userOption < min || userOption > max {
			fmt.Printf("Option not within valid range (%d - %d). Try again!\n", min, max)
			continue
		}

		break
	}

	return userOption
}

// getValidRange returns the valid range for user options
func getValidRange(section *Section) (min, max int) {
	return 0, len(section.NumberedChildren)
}

func userChoseExit(option int) bool {
	return option == exitOptionValue
}

func userChoseBack(option int, section *Section) bool {
	return (section.NumberedChildren)[option] == backOptionName
}
