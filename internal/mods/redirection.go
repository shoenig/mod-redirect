package mods

import (
	"encoding/json"
	"strings"
)

// Redirection is the representation of a go-get redirect that should be
// observed by a requesting entity (e.g. the Go compiler).
type Redirection struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	VCS         string `json:"vcs"`
	Destination string `json:"destination"`
}

func Key(kind, module string) []byte {
	return []byte(strings.Join([]string{
		kind, module,
	}, ","))
}

func (r *Redirection) Key() []byte {
	return Key(r.Kind, r.Name)
}

func (r *Redirection) Bytes() []byte {
	bs, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return bs
}
