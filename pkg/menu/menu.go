package menu

import (
	"fmt"
	"reflect"
)

// MenuChildren
// TODO: change its name
// TODO: implement string method for this type
type MenuChildren map[string]Menuer

// MetaData contains the minimum data for a Section or Exercise
type MetaData struct {
	Id          string // Item identifier
	Description string // Item description
}

// Section contains metadata and children bound to it
type Section struct {
	MetaData
	Children MenuChildren
}

// Exercise contains metadata and a function that implements the exercise's code
type Exercise struct {
	MetaData
	Runner func(args ...interface{}) error
}

// Menuer is the interface related to the MenuChildren type
// TODO: Change its name
type Menuer interface {
	String() string
	Name() string
	Desc() string
	Attach(menuer Menuer) error
}

// Options is used as the argument type for the Get function,
// since the From argument can be optional
type Options struct {
	Name string
	From Menuer
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

// Attach doesn't work on SingleSections
func (ss *MetaData) Attach(menuer Menuer) error {
	panic("single sections are not allowed to have children")
}

// String returns the string representation of a Section
func (s *Section) String() string {
	return fmt.Sprintf("<Name: %s> <Description: %s> <Children: %s>", s.Id, s.Description, s.Children)
}

// Name returns the Section's name
func (s *Section) Name() string {
	return s.Id
}

func (s *Section) Desc() string {
	return s.Description
}

// Attach binds a MetaData, Section or Exercise to a Section
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

// String returns the string representation of as Exercise
func (e *Exercise) String() string {
	return fmt.Sprintf("%s", e.MetaData)
}

// Name returns the Exercise's name
func (e *Exercise) Name() string {
	return e.Id
}

func (e *Exercise) Desc() string {
	return e.Description
}

// Attach doesn't work on Exercises
func (e *Exercise) Attach(menuer Menuer) error {
	panic("exercises are not allowed to have children")
}

var topMenu Section

// Add adds an item of name 'name' to another item of type Menuer
func Add(name string, item Menuer) {

	if isMenuEmpty() {
		topMenu.Id = "GoGym"
		topMenu.Description = "Exercising in Go"
		topMenu.Children = make(MenuChildren)
	}

	if _, ok := topMenu.Children[name]; ok {
		panic("item already in menu")
	}

	topMenu.Children[name] = item
}

// Get returns an element of type Menuer that has already been added to the menu, or error if the item is not found
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

// isMenuEmpty returns true if no children have yet been attached to the topMenu variable
func isMenuEmpty() bool {
	return topMenu.Children == nil
}

// equal returns true if two variables of the Menuer interface are equal, false otherwise
func equal(a, b Menuer) bool {
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
	return a.Id == b.Id &&
		a.Description == b.Description
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
	return a.Id == b.Id &&
		a.Description == b.Description
	//TODO: Compare MenuChildren
	// &&
	// a.Children == b.Children
}

type buildOptions struct {
	from Section
}

// Implements the main loop for gogym
func Loop() error {
	var buildOps buildOptions
	var option, min, max int
	var tempMenu *map[int]string

	for {
		tempMenu = buildMenu(&buildOps)
		display(tempMenu, &buildOps)
		fmt.Println("select an option from the menu...")
		i, err := fmt.Scanf("%d", &option)

		fmt.Printf("%d, %v\n", i, err)
		min, max = getValidRange(tempMenu)
		if option < min || option > max {
			fmt.Printf("option not within valid range (%d - %d). Try again!\n", min, max)
			continue
		}

		//TODO: option 0 should be a constant or something
		if option == 0 {
			fmt.Println("Bye!")
			break
		}

		//TODO: resolve choosen Menuer underlaying type and build another menu or run function

	}

	return nil
}

// display shows the menu to the user
func display(menu *map[int]string, options *buildOptions) {
	var navigator Section

	if options.from.Children == nil {
		navigator = topMenu

	} else {
		navigator = options.from
	}

	for k, v := range *menu {
		fmt.Printf("%d.%20s%20s\n", k, v, navigator.Children[v].Desc())
	}
}

//TODO: Change return type to numberedMenu or similar
func buildMenu(options *buildOptions) *map[int]string {
	var temp = make(map[int]string)
	var navigator Section

	if options.from.Children == nil {
		navigator = topMenu

	} else {
		navigator = options.from
	}

	i := 1
	for k := range navigator.Children {
		temp[i] = k
		i++
	}

	temp[i] = "Exit"

	return &temp
}

func getValidRange(menu *map[int]string) (min, max int) {
	size := len(*menu)
	if size == 0 {
		return 0, 0
	}

	return 1, size

}
