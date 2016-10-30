package gitdeps
// vim: noet ts=4 sw=4 sr smartindent:

import (
    "fmt"

	toml "github.com/BurntSushi/toml"
)

type TomlCfg struct {
	Deps map[string]DepInfo
}

func Read(toml_file string) (TomlCfg, error) {

	var c TomlCfg

	_, err := toml.DecodeFile(toml_file, &c)
	if err != nil {
		err_msg := "ERROR: [%s] - problem with toml file:\n%s\n"
		err = fmt.Errorf(err_msg, toml_file, err)
		return c, err
	}

	err = c.ValidateDeps(toml_file)
	return c, err
}

func (c TomlCfg) ValidateDeps(toml_file string) (err error) {

	fmt.Printf("INFO: [%s] - validating toml file\n", toml_file)

	for clone_dir, d := range c.Deps {
		err = d.Validate(toml_file, clone_dir)
		if err != nil {
			return err
		}
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
