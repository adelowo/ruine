package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const ruinableFileExtension = ".go"

var (
	dir   string
	files []string
	re    *regexp.Regexp
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

	if isGitRepo(dir) {
		defer ruin()
	}

	fmt.Printf("%d go source files were successfully ruined \n", len(files))
}

func pathWalker(path string, file os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("An error occurred while trying to walk the directory, %s. Err := %v \n", dir, err)
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
	return path[len(path)-3:] == ruinableFileExtension
}

func isGitRepo(dir string) bool {

	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	_, err := ioutil.ReadDir(dir + ".git")

	return err == nil
}

func ruin() {
	os.RemoveAll(dir + ".git")
	os.Mkdir("dir"+".git", os.ModeDir)

	out, err := exec.Command("git", "init").Output()

	if err != nil {
		fmt.Println("We couldn't complete the ruination")
		os.Exit(1)
	}

	if !strings.Contains(string(out), "Initialized empty Git repository") {
		fmt.Println("Git failed us")
		os.Exit(1)
	}
}
