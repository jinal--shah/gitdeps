package gitdeps
// vim: et ts=4 sw=4 sr smartindent:
import (
    "fmt"
)

type Gitdeps struct {
    All []Gitdep
}

type Gitdep struct {
    File      string
    Section   string
    *DepInfo
}

