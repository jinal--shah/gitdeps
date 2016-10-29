package toml_cfg
// vim: noet ts=4 sw=4 sr smartindent:

import (
	"errors"
    "fmt"

	toml "github.com/BurntSushi/toml"
)

type TomlCfg struct {
	Deps map[string]DepInfo
}

type DepInfo struct {
	Src   string
	Ref   string
	Depth string
}

func Read(toml_file string) (TomlCfg, error) {

	var c TomlCfg

	_, err := toml.DecodeFile(toml_file, &c)
	if err != nil {
		fmt.Println(err)
		return c, err
	}

	err = c.ValidateDeps(toml_file)
	return c, err
}

func (c TomlCfg) ValidateDeps(toml_file string) error {

	var err error

	fmt.Println("INFO: validating toml file " + toml_file)

	for clone_dir, d := range c.Deps {
		err = d.ValidateRef(clone_dir)
		if err != nil {
			return err
		}
	}

	return err
}

func (d DepInfo) ValidateRef(clone_dir string) error {

	var err error

	if d.Ref == "" {
		err_tmpl := "ERROR: ref not supplied for src %s in [deps.%s]\n"
		err_msg := fmt.Sprintf(err_tmpl, d.Src, clone_dir)
		err = errors.New(err_msg)
	}

	return err
}

/* 

    We expect a toml file such as this:

	[deps]
		[deps.<dir to clone in to>]
		src   = "<any uri accepted by git clone cmd>"
		ref   = "<optional value accepted by --branch option>"
		depth = "<optional value accepted by --depth option>"

		[deps.build_alpine]
		src = "git@github.com:EurostarDigital/build_ami"
		ref = "master"
		depth = "1"

		[deps.coreos_setup]
		src = "git://github.com/jinal--shah/demo-coreos-vagrant-setup"
		ref = "add-docker-cleanup-script"

*/
