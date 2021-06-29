package menu

import (
	"bufio"
	"fmt"
	"path"
)

// Implements the main loop for gogym
func Loop(root string) error {

	var userOption, whereUserIs string = "", root
	var reader *bufio.Reader

	for {
		fmt.Printf("%s %s%s\n", "You are in", whereUserIs, ". The menu is:")
		retrieveMenu(userOption)
		userOption, _ = reader.ReadString('\n')
		whereUserIs = path.Join(whereUserIs, getUserLocation(userOption))
	}

}

// Reads the corresponding directory looking for its menu
func retrieveMenu(option string) {

}

// Gets the user's current location
func getUserLocation(option string) string {
	return ""
}
