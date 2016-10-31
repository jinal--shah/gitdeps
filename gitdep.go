package gitdeps
// vim: et ts=4 sw=4 sr smartindent:
import (
    "fmt"
    "strconv"
    "strings"
    "path/filepath"
)

type Gitdep struct {
    File        string      // abs path to .gitdeps file
    CloneDir    string      // dir within .gitdeps dir to clone in to
    Src         string      // toml: git origin location
    Ref         string      // toml: git clone / checkout arg
    Depth       string      // toml: git clone / checkout arg
    Errs                   // promoted methods for error aggregation and visibility
}

func (g *Gitdep) e_context(msg string) (string) {
    return fmt.Sprintf("ERROR: [file:%s][deps.%s]: %s\n", g.File, g.CloneDir, msg)
}

// ... add file and clone_dir info to Gitdep, and validate
func (g *Gitdep) Configure(toml_file string, clone_dir string) (err error) {
    g.File     = toml_file
    g.CloneDir = clone_dir

    err        = g.Validate()
    return err
}

func (g *Gitdep) Validate() (err error) {
    g.ValidateCloneDir()
    g.ValidateSrc()
    g.ValidateRef()
    g.ValidateDepth()
    err = g.sprintfe(g) // err only has a value if any err msgs captured
    return err
}

func (g *Gitdep) ValidateCloneDir() {
    if strings.Contains(g.CloneDir, "/") {
        g.e("clone path can not be multiple directories")
    }
}

func (g *Gitdep) ValidateSrc() {
    if g.Src == "" {
        g.e("src not supplied")
    }
}

func (g *Gitdep) ValidateRef() {
    if g.Ref == "" {
        g.e("ref not supplied")
    }
}

// Depth: must be int if defined, non-empty
func (g *Gitdep) ValidateDepth() {
    if g.Depth != "" {
        if _, err := strconv.Atoi(g.Depth); err != nil {
            g.e("depth is not a valid number - provided: " + g.Depth)
        }
    }
}

func (g *Gitdep) ClonePath() (string) {
    return filepath.Dir(g.File) + "/" + g.CloneDir
}

