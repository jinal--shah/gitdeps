package gitdeps
// vim: et ts=4 sw=4 sr smartindent:

import (
    "path/filepath"
    "os"
    "fmt"
)

const filename_match = ".gitdeps"

// array of abs paths matching filename
type Files struct {
    StartDir  string
    Found []string
    Console
}

func NewFiles(start_dir string) *Files {
    return &Files{StartDir: start_dir}
}

// satisfy context interface
func (g *Files) context() (string) {
    return fmt.Sprintf("[start_dir:%s][finding %s]", g.StartDir, filename_match)
}

func (g *Files) filterMatches(path string, f os.FileInfo, err error) (error) {
    name    := filepath.Base(path)
    ok, _   := filepath.Match(filename_match, name)
    if ok {
        path, _   = filepath.Abs(path)
        g.Found = append(g.Found, path)
    }
    return err
}

func (g *Files) Recursively() (file_list []string, err error) {

    err = filepath.Walk(g.StartDir, g.filterMatches)

    if err != nil {
        g.e(err.Error())
    }

    file_list = g.Found
    err = g.sprintfe(g)

    if len(file_list) > 0 {
        for _, file := range file_list {
            g.printfi(g, "found " + file)
        }
    } else {
        g.printfi(g, "No matching files found.")
    }

    return file_list, err
}

