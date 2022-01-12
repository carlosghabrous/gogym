package bluebook

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func usage(progName string) error {
	return fmt.Errorf("usage: %s <file.m3u>", progName)
}

func Exercise01C03(args ...interface{}) error {
	if len(os.Args) != 1 || !strings.HasSuffix(os.Args[1], ".m3u") {
		return usage(os.Args[0])
	}

	if content, err := ioutil.ReadFile(os.Args[1]); err != nil {
		return err

	} else {

		songs := readSongsFromM3uFile(content)
		writeSongsToPlsFile(songs)
	}

	return nil
}

func readSongsFromM3uFile(content []byte) []string {
	return []string{}
}

func writeSongsToPlsFile(songs []string) {

}
