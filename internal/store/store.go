package store

import (
	"fmt"

	"gophers.dev/cmds/mod-redirect/internal/mods"
)

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -g -i Storage -s _mock.go

type Storage interface {
	// Get redirection for (kind, namespace (path))
	Get(string, string) (*mods.Redirection, error)

	// List all module redirections
	List() mods.Redirects
}

func New(redirects mods.Redirects) Storage {
	return &embedStore{
		modules: redirects,
	}
}

type embedStore struct {
	modules mods.Redirects
}

func (es *embedStore) Get(kind, namespace string) (*mods.Redirection, error) {
	// todo: use a multi map instead

	for _, m := range es.modules {
		if m.Kind == kind && m.Namespace == namespace {
			return m.Copy(), nil
		}
	}

	return nil, fmt.Errorf("module %s/%s not found", kind, namespace)
}

func (es *embedStore) List() mods.Redirects {
	return es.modules.Copy()
}
