package menu

import (
	"fmt"
	"reflect"
	"strings"
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

var topMenu Section

// Add adds an item of name 'name' to another item of type MenuItemer
func Add(name string, item nameDescriptioner) {

	if isMenuEmpty() {
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
	var buildOps buildOptions
	var option, min, max int
	var tempMenu *numberedMenu

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

		//TODO: resolve choosen MenuItemer underlaying type and build another menu or run function

	}

	return nil
}

// options is used as the argument type for the get function,
// since the from argument can be optional
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
		fmt.Errorf("only Section(s) contain children elements")
	}

	returnMenuer, ok = fromSection.Children[options.name]
	if !ok {
		return returnMenuer, fmt.Errorf("top menu doesn't contain item %s", options.name)
	}

	return returnMenuer, nil
}

// isMenuEmpty returns true if no children have yet been attached to the topMenu variable
func isMenuEmpty() bool {
	return topMenu.Children == nil
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

// display shows the menu to the user
func display(menu *numberedMenu, options *buildOptions) {
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

func buildMenu(options *buildOptions) *numberedMenu {
	var temp = make(numberedMenu)
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

func getValidRange(menu *numberedMenu) (min, max int) {
	size := len(*menu)
	if size == 0 {
		return 0, 0
	}

	return 1, size

}
