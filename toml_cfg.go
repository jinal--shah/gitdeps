package gitdeps
// vim: et ts=4 sw=4 sr smartindent:

import (
    "fmt"

    toml "github.com/jinal--shah/toml"
)

type TomlCfg struct {
    Gitdeps  map[string]*Gitdep // pointers to Gitdep lets us modify in for loops
}

func Read(toml_file string) (c TomlCfg, err error) {

    _, err = toml.DecodeFile(toml_file, &c)
    if err != nil {
        err_msg := "ERROR: [%s] - problem with toml file:\n%s\n"
        err = fmt.Errorf(err_msg, toml_file, err)
        return c, err
    }

    for clone_dir, g := range c.Gitdeps {
        err = g.Configure(toml_file, clone_dir)
        if err != nil {
            return c, err
        }
    }

    return c, err
}

/* 

    We expect a toml file such as this:

    [gitdeps]
        [gitdeps.<dir to clone in to>]
        src   = "<any uri accepted by git clone cmd>"
        ref   = "<optional value accepted by --branch option>"
        depth = "<optional value accepted by --depth option>"

        [gitdeps.build_alpine]
        src = "git@github.com:EurostarDigital/build_ami"
        ref = "master"
        depth = "1"

        [gitdeps.coreos_setup]
        src = "git://github.com/jinal--shah/demo-coreos-vagrant-setup"
        ref = "add-docker-cleanup-script"

*/
