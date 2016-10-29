package find_files
// vim: noet ts=4 sw=4 sr smartindent:

import (
	"path/filepath"
	"os"
	"fmt"
)

// array of abs paths matching filename
var file_list []string

// filename to match
var filename_match string

func filterMatches(path string, f os.FileInfo, err error) error {
	name    := filepath.Base(path)
	ok, _   := filepath.Match(filename_match, name)
	if ok {
		path, _   = filepath.Abs(path)
		file_list = append(file_list, path)
	}
	return nil
}

func Recursively(root string, f string) ([]string, error) {
	filename_match = f

	err := filepath.Walk(root, filterMatches)

	if err != nil {
		fmt.Printf("ERROR: ", err)
	}
	return file_list, err
}

