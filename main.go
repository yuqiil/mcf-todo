package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var root string

	// Get the path of the current directory
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Call to function to browse through javascript files
	files, err := FilePathWalkDir(root)
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}

func FilePathWalkDir(root string) (string, error) {
	var finallist string
	var files []string

	// Regexp for exact word 'TODO'
	var regextodo = regexp.MustCompile(`\bTODO\b`)

	// Walks a file tree for each file or directory in the tree, including the root
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		// Check that files are of javascript type
		if filepath.Ext(info.Name()) == ".js" {

			// Read file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			// Convert file content into string
			og := string(content)

			// Check for 'TODO's in the file. Add path to list for printing afterwards if TODO is found.
			result := regextodo.MatchString(og)
			if result == true {
				files = append(files, path)
			}
		}
		return nil
	})
	// Convert bytes to strings to produce the final list of files with TODOs
	finallist = strings.Join(files, "\n")
	return finallist, err
}
