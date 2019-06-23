package service

import (
	"net/http"
	"os"

	"github.com/shoenig/loggy"
	"github.com/shoenig/mod-redirect/config"
	"github.com/shoenig/mod-redirect/internal/store"
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
			r.log.Errorf("faield to initialize proxy")
			os.Exit(1)
		}
	}

	return r
}
