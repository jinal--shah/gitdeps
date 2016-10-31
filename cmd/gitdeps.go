package main
// vim: et ts=4 sw=4 sr smartindent:

import (
    gd "github.com/jinal--shah/gitdeps"

    "fmt"
    "os"
)

func main() {
    root_dir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    err = ProcessDepsFiles(root_dir)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}

func ProcessDepsFiles(start_dir string) (err error) {
    f := gd.NewFiles(start_dir)
    file_list, err := f.Recursively()
    if err != nil {
        fmt.Println(err)
        return err
    }

    if len(file_list) == 0 {
        fmt.Printf("INFO: [start_dir:%s] - No files found\n", start_dir)
    }

    for _, file_name := range file_list {
        err = ProcessFile(file_name)
    }
    return err

}

func ProcessFile(file_name string) (err error) {
    c, err := gd.Read(file_name)
    if err != nil {
        return err
    }

    for _, g := range c.Gitdeps {
        _, err = g.GitClone();
        if err != nil {
            fmt.Println(err)
        }
    }

    return err
}

