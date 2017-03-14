package main

import (
	"flag"
	"os"
	"path/filepath"
	"fmt"
	"regexp"
	"io/ioutil"
)

const ruinableFileExtension = ".go"

var (
	dir string
	files []string
	re *regexp.Regexp
)

func init() {
	flag.StringVar(&dir, "dir", "./", "The directory in which go files are to be ruined")
	re = regexp.MustCompile(`error|err|\.`)
}

func main() {

	flag.Parse()

	filepath.Walk(dir, pathWalker)

	for _, v := range files {
		contents, err := ioutil.ReadFile(v)
		check(err)
		newContents := re.ReplaceAllString(string(contents), "")
		check(ioutil.WriteFile(v, []byte(newContents), 0644))
	}

	fmt.Printf("%d go source files were successfully ruined \n", len(files))
}

func pathWalker(path string, file os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("An error occurred while trying to walk the directory, %s. Err := %v", dir, err)
		//os.Exit(1)
		return nil
	}

	if !file.IsDir() && isGoFile(path) {
		files = append(files, path)
	}

	return nil
}

func check(err error) {
	if err != nil {
		fmt.Printf("An error occured %v", err)
		os.Exit(1)
	}
}

func isGoFile(path string) bool {
	return path[len(path) - 3 :] == ruinableFileExtension
}
