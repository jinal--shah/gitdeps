package main
// package main: to generate binary `gitdeps`
// vim: et ts=4 sw=4 sr smartindent:

import (
    gd "github.com/jinal--shah/gitdeps"

    "fmt"
    "os"
)

func main() {

    cmd, err := MakeCmd()

    if err != nil {
        os.Exit(1)
    }

    err = ProcessDepsFiles(&cmd)
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

