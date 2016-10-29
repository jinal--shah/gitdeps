package dep
// vim: noet ts=4 sw=4 sr smartindent:
import (
    "fmt"
)

type DepInfo struct {
	Src   string
	Ref   string
	Depth string
}

func (d DepInfo) Validate(toml_file string, clone_dir string) (err error) {

	failed := false

	err = d.ValidateSrc(toml_file, clone_dir)
	if err != nil {
		failed = true
	}
	err = d.ValidateRef(toml_file, clone_dir)
	if err != nil {
		failed = true
	}

	if failed {
		err_tmpl := "ERROR: [%s] can't continue with [deps.%s]\n"
		err = fmt.Errorf(err_tmpl, toml_file, clone_dir)
	}

	return err
}

func (d DepInfo) ValidateSrc(toml_file string, clone_dir string) (err error) {

	if d.Src == "" {
		err_tmpl := "ERROR: [%s] src not supplied for [deps.%s]\n"
		err = fmt.Errorf(err_tmpl, toml_file, clone_dir)
	}

	return err
}

func (d DepInfo) ValidateRef(toml_file string, clone_dir string) (err error) {

	if d.Ref == "" {
		err_tmpl := "ERROR: [%s] ref not supplied for [deps.%s]\n"
		err = fmt.Errorf(err_tmpl, toml_file, clone_dir)
	}

	return err
}
