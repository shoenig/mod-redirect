package mods

import (
	"encoding/json"
	"errors"
	"strings"
)

type Redirects []*Redirection

func (rs Redirects) Copy() Redirects {
	c := make(Redirects, len(rs))
	for i := 0; i < len(rs); i++ {
		c[i] = rs[i].Copy()
	}
	return c
}

// Redirection is the representation of a go-get redirect that should be
// observed by a requesting entity (e.g. the Go compiler).
type Redirection struct {
	Kind        string `json:"kind"`        // i.e. cmds or pkgs
	Namespace   string `json:"namespace"`   // e.g. foo/bar/baz
	VCS         string `json:"vcs"`         // i.e. git, https, hg
	Destination string `json:"destination"` // e.g. github.com/foo/bar
}

var (
	ErrMissingKind        = errors.New("kind field must not be empty")
	ErrMissingNamespace   = errors.New("namespace field must not be empty")
	ErrMissingVCS         = errors.New("vcs field must not be empty")
	ErrMissingDestination = errors.New("destination field must not be empty")
	ErrMissingProtocol    = errors.New("protocol in destination is required")
)

func Valid(r *Redirection) error {
	if r.Kind == "" {
		return ErrMissingKind
	}

	if r.Namespace == "" {
		return ErrMissingNamespace
	}

	if r.VCS == "" {
		return ErrMissingVCS
	}

	if r.Destination == "" {
		return ErrMissingDestination
	}

	if !strings.Contains(r.Destination, "://") {
		return ErrMissingProtocol
	}

	return nil
}

func (r *Redirection) String() string {
	return r.Destination
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

func (r *Redirection) Copy() *Redirection {
	return &Redirection{
		Kind:        r.Kind,
		Namespace:   r.Namespace,
		VCS:         r.VCS,
		Destination: r.Destination,
	}
}
