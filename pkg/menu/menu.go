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

// MetaData contains the minimum data for a Section or Exercise
type MetaData struct {
	Id          string // Item identifier
	Description string // Item description
}

// Section contains metadata and children bound to it
type Section struct {
	MD       MetaData
	Children offspring
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

type numberedMenu map[int]string

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
func (s *Section) Attach(item nameDescriptioner) error {
	if s.Children == nil {
		s.Children = make(offspring)
	}

	name := item.Name()

	if _, ok := s.Children[name]; ok {
		return fmt.Errorf("item %s already exists", name)
	}

	s.Children[name] = item
	return nil
}

// String returns the string representation of as Exercise
func (e *Exercise) String() string {
	return fmt.Sprintf("%s", e.MD)
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
	}

	if _, ok := topMenu.Children[name]; ok {
		panic("item already in menu")
	}

	topMenu.Children[name] = item
}

// Implements the main loop for gogym
func Loop() error {
	var userOption int
	var tempMenu *numberedMenu
	var menuStack *stack.Stack = stack.New()
	var theChosenOne interface{}
	var buildOps = buildOptions{from: topMenu}

	for {
		tempMenu = buildNumberedMenu(&buildOps)

		display(tempMenu, &buildOps)

		userOption = getUserOption(tempMenu)

		if userChoseExit(userOption, tempMenu) {
			break
		}

		if userChoseBack(userOption, tempMenu) {
			theChosenOne = menuStack.Pop()

		} else {
			theChosenOne = buildOps.from.Children[(*tempMenu)[userOption]]
			menuStack.Push(buildOps.from)
		}

		switch theChosenOne := theChosenOne.(type) {

		case *Exercise:
			theChosenOne.Runner()

		case *Section:
			buildOps.from = *theChosenOne
		}
	}

	return nil
}

// options is used as the argument type for the get function,
// since the from argument can be optional
func addExtraOption(s *Section, numMenu *numberedMenu, options *menuOptionTuple) {
	s.Children[options.name] = &Section{MD: MetaData{Id: options.name, Description: options.description}}
	(*numMenu)[options.value] = options.name
}

type options struct {
	name string
	from nameDescriptioner
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

type buildOptions struct {
	from Section
}

type menuOptionTuple struct {
	name        string
	description string
	value       int
}

func buildNumberedMenu(options *buildOptions) *numberedMenu {
	var temp = make(numberedMenu)

	// Add extra 'back' option if not in the top level menu
	if options.from.MD.Id != topMenu.MD.Id {
		addExtraOption(&options.from, &temp, &menuOptionTuple{backOptionName, backOptionDescription, len(options.from.Children) + 1})
	}

	i := 1
	for k := range options.from.Children {
		if k == backOptionName {
			continue
		}

		fmt.Printf("Adding %d - %v to the temp menu\n", i, k)
		temp[i] = k
		i++
	}

	addExtraOption(&options.from, &temp, &menuOptionTuple{exitOptionName, exitOptionDescription, exitOptionValue})
	return &temp
}

// display shows the menu to the user
func display(menu *numberedMenu, options *buildOptions) {
	var name, description, formatString string
	formatString = "-> %-10d%-20s%-20s\n"

	for i := 1; i < len(*menu); i++ {
		name = (*menu)[i]
		description = options.from.Children[name].Desc()
		fmt.Printf(formatString, i, name, description)
	}

	fmt.Printf(formatString, exitOptionValue, exitOptionName, exitOptionDescription)
}

func getUserOption(menu *numberedMenu) int {
	var userOption int

	for {

		fmt.Println("Select an option from the menu")

		_, err := fmt.Scanf("%d", &userOption)
		if err != nil {
			fmt.Println("Something went wrong! Try again!")
			continue
		}

		min, max := getValidRange(menu)
		if userOption < min || userOption > max {
			fmt.Printf("Option not within valid range (%d - %d). Try again!\n", min, max)
			continue
		}

		break
	}

	return userOption
}

// getValidRange returns the valid range for user options
func getValidRange(menu *numberedMenu) (min, max int) {
	return 0, len(*menu) - 1
}

func userChoseExit(option int, menu *numberedMenu) bool {
	return (*menu)[option] == exitOptionName
}

func userChoseBack(option int, menu *numberedMenu) bool {
	return (*menu)[option] == backOptionName
}
