package gitdeps
// vim: et ts=4 sw=4 sr smartindent:

import (
    "fmt"
    "os/exec"
    "strings"
)

const cmd = "git"

func (g *Gitdep) Options() (options []string) {

    if g.Ref != "" {
        options = append(options, "--branch", g.Ref)
    }

    if g.Depth != "" {
        options = append(options, "--depth", g.Depth)
    }

    return options
}

// returns array for ClondCmd
func (g *Gitdep) GitCloneCmdArgs() (clone_args []string) {

    options := g.Options()

    clone_args = append(clone_args, "clone", g.Src)
    if len(options) > 0 {
        clone_args = append(clone_args, options...)
    }
    clone_args = append(clone_args, g.ClonePath() )

    return clone_args
}

func (g *Gitdep) GitClone() ([]byte, error) {

    cmd_args := g.GitCloneCmdArgs()

    g.printfi(g, "... executing:")
    g.printfi(g, fmt.Sprintf("%s %s", cmd, strings.Join(cmd_args[:], " ")))

    out, err := exec.Command(cmd, cmd_args...).CombinedOutput()

    if err != nil {
        g.e(string(out))
        g.e(err.Error())
    } else {
        g.printfi(g, string(out))
    }

    err = g.sprintfe(g)
    return out, err
}

