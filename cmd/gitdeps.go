package main
// vim: et ts=4 sw=4 sr smartindent:

import (
    gd "github.com/jinal--shah/gitdeps"

    "path/filepath"
    "fmt"
    "os"
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
    file_list, err := gd.Recursively(start_dir, filename_match)
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
    c, err := gd.Read(file_name)
    if err != nil {
        os.Exit(1)
    }

    clone_path_base := filepath.Dir(file_name)

    for dest_dir, dep := range c.Deps {
        cmd_args := gd.GitCloneCmdArgs(clone_path_base + "/" + dest_dir, dep)
        _, _ = gd.RunClone(cmd_args);
    }

    return err
}

