package main
// vim: noet ts=4 sw=4 sr smartindent:

import (
	"find_files"
	"toml_cfg"
	git "git_functions"

	"path/filepath"
	"fmt"
	"os"
)

type Gitdep struct {
	*dep.DepInfo
	*git.GitCmds
}

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
		os.Exit(1)
	}

	if len(file_list) > 0 {
		fmt.Printf("INFO: [looking under %s] - found files:\n%v\n", start_dir, file_list)
	} else {
		fmt.Printf("INFO: [looking under %s] - No files found\n", start_dir)
	}

	for _, file_name := range file_list {
		err = ProcessFile(file_name)
	}

}

func ProcessFile(file_name string) error {
	c, err := toml_cfg.Read(file_name)
	if err != nil {
		os.Exit(1)
	}

	clone_path_base := filepath.Dir(file_name)

	for dest_dir, dep := range c.Deps {
		cmd_args := git.GitCloneCmdArgs(clone_path_base + "/" + dest_dir, dep)
		_, _ = git.RunClone(cmd_args);
	}

	return err
}

