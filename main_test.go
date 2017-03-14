package main

import (
	"testing"
	"path/filepath"
)

func TestIsGoFile(t *testing.T) {

	table := []struct {
		path string
		valid bool
	} {
		{"gofile.go", true},
		{"gofile.py", false},
		{"gofile.php", false},
		{"gofile.rb", false},
	}

	for _,v := range table {
		if got := isGoFile(v.path); got != v.valid {
			t.Errorf("Want %v. Got %v", v.valid, got)
		}
	}
}


func TestPathWalker(t *testing.T) {

	filepath.Walk(dir, pathWalker)

	//Fragile test though
	//2 since pathWalker internally acknowledges *.go files only
	if len(files) != 2 {
		t.Errorf("This project currently has %d files. Got %d instead", 2, len(files))
	}
}
