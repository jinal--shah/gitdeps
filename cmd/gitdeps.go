package main
// vim: et ts=4 sw=4 sr smartindent:

import (
    gd "github.com/jinal--shah/gitdeps"

    "fmt"
    "os"
)

// we always begin from current working directory
var err_msg string

func main() {
    root_dir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }

    ProcessDepsFiles(root_dir)

}

func ProcessDepsFiles(start_dir string) {
    file_list, err := gd.NewFiles(start_dir).Recursively()
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
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("... starting cloning for " + file_name)
    for _, g := range c.Gitdeps {
        _, err = g.GitClone();
    }

    return err
}

