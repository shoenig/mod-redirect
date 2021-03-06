package service

import (
	"net/http"
	"os"

	"gophers.dev/cmds/mod-redirect/config"
	"gophers.dev/cmds/mod-redirect/internal/store"

	"gophers.dev/pkgs/loggy"
)

type Redirect struct {
	config  config.Configuration
	storage store.Storage
	server  *http.ServeMux
	log     loggy.Logger
}

func NewRedirect(configuration config.Configuration) *Redirect {
	r := &Redirect{
		config: configuration,
		log:    loggy.New("mod-redirect"),
	}

	for _, f := range []initer{
		initStore,
		initWeb,
	} {
		if err := f(r); err != nil {
			r.log.Errorf("failed to initialize proxy: %v", err)
			os.Exit(1)
		}
	}

	return r
}

// Start the thing.
func (r *Redirect) Start() {
	r.log.Infof("--- starting! ---")
	// the web server is already running
	select {
	// intentionally left blank
	}
}
