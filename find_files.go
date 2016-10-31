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
}

func NewFiles(start_dir string) *Files {
    return &Files{StartDir: start_dir}
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
        err = fmt.Errorf("ERROR: [looking under %s] %s", g.StartDir, err)
    }
    file_list = g.Found
    return file_list, err
}

