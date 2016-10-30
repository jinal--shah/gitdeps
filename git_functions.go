package gitdeps
// vim: et ts=4 sw=4 sr smartindent:

import (
    "fmt"
    "os/exec"
    "strings"
)

// git clone <d.Src> <opt_str> <clone_path>
const clone_cmd_tmpl = "git clone %s %s %s"
const cmd = "git"

func Options(d DepInfo) []string {

    var options []string

    if d.Ref != "" {
        options = append(options, "--branch", d.Ref)
    }

    if d.Depth != "" {
        options = append(options, "--depth", d.Depth)
    }

    return options
}

// returns array for Exec.Cmd.Run()
func GitCloneCmdArgs(clone_path string, d DepInfo) []string {

    var clone_cmd []string

    options := Options(d)

    clone_cmd = append(clone_cmd, "clone", d.Src)
    if len(options) > 0 {
        clone_cmd = append(clone_cmd, options...)
    }
    clone_cmd = append(clone_cmd, clone_path)

    return clone_cmd
}

func RunClone(cmd_args []string) ([]byte, error) {
    fmt.Printf("%s %s\n", cmd, strings.Join(cmd_args[:], " "))

    out, err := exec.Command(cmd, cmd_args...).CombinedOutput()

    if err != nil {
        fmt.Printf("%s\n", err)
    }
    fmt.Printf("%s\n", out)

    return out, err
}

