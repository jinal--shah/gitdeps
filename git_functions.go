package gitdeps
// vim: et ts=4 sw=4 sr smartindent:

import (
    "fmt"
    "os"
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

// returns slice for GitClone
func (g *Gitdep) GitCloneCmdArgs() (clone_args []string) {

    options := g.Options()

    clone_args = append(clone_args, "clone", g.Src)
    if len(options) > 0 {
        clone_args = append(clone_args, options...)
    }
    clone_args = append(clone_args, g.ClonePath() )

    return clone_args
}

func (g *Gitdep) GitClone() (out []byte, err error) {

    cmd_args := g.GitCloneCmdArgs()

    g.printfi(g, "... cloning:")
    out, err = g.GitCmd(cmd_args)
    if err != nil {
        g.e(string(out))
        g.e(err.Error())
    } else {
        g.printfi(g, string(out))
    }

    err = g.sprintfe(g)
    return out, err
}

func (g *Gitdep) GitCmd(cmd_args []string) ([]byte, error) {

    g.printfi(g, fmt.Sprintf("%s %s", cmd, strings.Join(cmd_args[:], " ")))

    return exec.Command(cmd, cmd_args...).CombinedOutput()

}

func (g *Gitdep) GitDir() (string) {
    return g.ClonePath() + "/.git"
}

// if not a gitrepo we must fail.
func (g *Gitdep) IsGitRepo() (bool) {
    cmd_args := []string{"--git-dir", g.GitDir(), "rev-parse"}

    g.printfi(g, "... testing if repo already cloned")
    out, err := g.GitCmd(cmd_args)
    if err != nil {
        g.printfi(g, "... Not a git repo.")
        return false
    } else {
        g.printfi(g, "... Is already a cloned git repo.")
        return true
    }
}

// We only run CurrentSrc on what we know is a gitrepo
// if not the expected src, we must fail.
func (g *Gitdep) CurrentSrc() (origin string, err error) {
    cmd_args := []string{"--git-dir", g.GitDir(), "config", "remote.origin.url"}

    g.printfi(g, "... determining repo's current origin")
    out, err := g.GitCmd(cmd_args)

    origin = string(out)
    if err != nil {
        g.e("output: " + origin)
        g.e(err.Error())
    } else {
        g.printfi(g, "output: " + origin)
    }

    err = g.sprintfe(g)
    return origin, err
}

func (g *Gitdep) IsExpectedSrc(origin string) (bool) {
    return origin == g.Src
}

func (g *Gitdep) CheckCurrentSrc() (err error) {
    origin, err := g.CurrentSrc()
    if err != nil {
        return err
    }

    if ! g.IsExpectedSrc(origin) {
        g.e("... current origin differs from requested src " + g.Src)
    }
    err = g.sprintfe(g)
    return err
}
