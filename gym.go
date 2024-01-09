package gym

import (
	"fmt"
)

type Gym struct {
	gym              map[string]*Section
	lastAddedSection int
}

// TODO: dummy -> move to its own file/package
type Section struct {
	name   string
	number int
}

func NewGym() *Gym {
	return &Gym{gym: make(map[string]*Section)}
}

func (g *Gym) AddSection(title string) (*Section, error) {
	if g.gym == nil {
		return nil, fmt.Errorf("create the gym before adding topics")
	}

	if g.itemExists(title) {
		return nil, fmt.Errorf("topic %s already exists", title)
	}

	section := &Section{name: title, number: g.lastAddedSection + 1}
	g.gym[title] = section
	g.lastAddedSection += 1
	return section, nil
}

func (g *Gym) itemExists(title string) bool {
	if _, found := g.gym[title]; found {
		return true
	}

	return false
}

// TODO: I might want to create an interface with this and other methods common
// to Gym, Section, Chapter and Exercise
func (g *Gym) Print() {
	for title, section := range g.gym {
		fmt.Println(title, section)
	}
}

// TODO: create another package or file with a 'runner'. This function
// should accept an interface (Gym, Section, Section or Exercise) and run it, controlling
// options, going back and forth, exiting...
