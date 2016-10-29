package main
// vim: noet ts=4 sw=4 sr smartindent:

import (
	"path/filepath"
	"find_files"
	"toml_cfg"
	"fmt"
	"os"

	git      "git_functions"
	readable "github.com/tonnerre/golang-pretty"
)

// we always begin from current working directory
var err_msg string
const filename_match = ".gitdeps"

func main() {
	root_dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	ProcessDepsFiles(root_dir)

}

func ProcessDepsFiles(start_dir string) {
	file_list, err := find_files.Recursively(start_dir, filename_match)
	if err != nil {
		panic("ERROR: ... couldn't recurse to find files:")
	} else {
		fmt.Printf("INFO: found files:\n%# v\n", readable.Formatter(file_list))
	}

	/* 	for each file in list
		- read toml
		- git clone
		- kick of ProcessDepsFiles for that dir
	*/

	for _, file_name := range file_list {
		err = ProcessFile(file_name)
	}

}

func ProcessFile(file_name string) error {
	c, err := toml_cfg.Read(file_name)
	if err != nil {
		err_msg = fmt.Sprintf("ERROR: ... couldn't read toml cfg\n%s", err)
		panic(err_msg)
	}

	clone_path_base := filepath.Dir(file_name)

	for dest_dir, dep := range c.Deps {
		cmd_args := git.GitCloneCmdArgs(clone_path_base + "/" + dest_dir, dep)
		_, _ = git.RunClone(cmd_args);
	}

	return err
}

