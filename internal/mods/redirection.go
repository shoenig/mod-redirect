package mods

import (
	"encoding/json"
	"strings"
)

// Redirection is the representation of a go-get redirect that should be
// observed by a requesting entity (e.g. the Go compiler).
type Redirection struct {
	Kind        string `json:"kind"`        // i.e. cmds or pkgs
	Namespace   string `json:"namespace"`   // e.g. foo/bar/baz
	VCS         string `json:"vcs"`         // i.e. git, https, hg
	Destination string `json:"destination"` // e.g. github.com/foo/bar
}

func Key(kind, module string) []byte {
	return []byte(strings.Join([]string{
		kind, module,
	}, ","))
}

func (r *Redirection) Key() []byte {
	return Key(r.Kind, r.Namespace)
}

func (r *Redirection) Bytes() []byte {
	bs, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return bs
}
