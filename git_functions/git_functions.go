package git_functions
// vim: noet ts=4 sw=4 sr smartindent:

import (
	"bytes"
	"fmt"
	"toml_cfg"
	"os/exec"
	"sync"
)

// git clone <d.Src> <opt_str> <clone_path>
const clone_cmd_tmpl = "git clone %s %s %s"

func Options(d toml_cfg.DepInfo) []string {

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
func GitCloneCmdArgs(clone_path string, d toml_cfg.DepInfo) []string {

	var clone_cmd []string

    options := Options(d)

	clone_cmd = append(clone_cmd, "clone", d.Src)
	if len(options) > 0 {
		clone_cmd = append(clone_cmd, options...)
	}
	clone_cmd = append(clone_cmd, clone_path)

	return clone_cmd
}

func RunClone(clone_cmd []string) string{
	fmt.Printf("INFO: running ... %v\n\n",clone_cmd)
}

func exe_cmd(cmd string, cmd_args []string, wg *sync.WaitGroup) {
	fmt.Println("command is ",cmd)
	// splitting head => g++ parts => rest of the command

	out, err := exec.Command(cmd, cmd_args...).CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}

// Given the clone_path exists, make sure it is the expected cloned repo
// ... doesn't check depth, just repo src and branch or tag
func IsExpectedRepo(clone_path string, d toml_cfg.DepInfo) (bool, error) {
	var err error
	return true, err
}
