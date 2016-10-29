package dep
// vim: noet ts=4 sw=4 sr smartindent:
import (
	"errors"
    "fmt"
)

type DepInfo struct {
	Src   string
	Ref   string
	Depth string
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

