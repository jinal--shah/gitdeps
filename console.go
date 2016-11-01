package gitdeps
// vim: et ts=4 sw=4 sr smartindent:
import (
    "fmt"
)

type Console struct {
    Errors []string
}

type console interface {
    context() string
}

func (c *Console) e(msg string) {
    c.Errors = append(c.Errors, msg)
}

// print INFO
func (c *Console) printfi(ci console, msg string) {
    context_msg := ci.context()
    fmt.Printf("INFO: %s: %s\n", context_msg, msg)
}

// print WARNING
func (c *Console) printfw(ci console, msg string) {
    context_msg := ci.context()
    fmt.Printf("WARNING: %s: %s\n", context_msg, msg)
}

func (c *Console) failed() (bool) {
    return len(c.Errors) > 0
}

// ERROR: doesn't print error msgs, just formats them.
func (c *Console) sprintfe(ci console) (err error) {
    if c.failed() {
        msg_combined := ""
        context_msg := ci.context()
        for _, msg := range c.Errors {
            msg = fmt.Sprintf("ERROR: %s: %s\n", context_msg, msg)
            msg_combined = msg + msg_combined
        }
        err = fmt.Errorf("%s", msg_combined)
    }

    return err
}

