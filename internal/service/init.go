package service

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pkg/errors"

	"gophers.dev/cmds/mod-redirect/internal/store"
	"gophers.dev/cmds/mod-redirect/internal/web"
)

type initer func(*Redirect) error

func initStore(r *Redirect) error {
	r.storage = store.New(r.config.Modules)
	r.log.Tracef("init store with %d modules", len(r.storage.List()))
	return nil
}

func initWeb(r *Redirect) error {
	r.log.Tracef("setting up web server @ %s", r.config.WebServer.Address())

	domain := r.config.Domain
	if domain == "" {
		return errors.New("domain must be specified")
	}

	router := mux.NewRouter()
	web.Set(router, domain, r.storage)

	server, err := r.config.WebServer.Server(router)
	if err != nil {
		return errors.Wrap(err, "could not create web server")
	}

	go func(h http.Handler) {
		panic(server.ListenAndServe())
	}(router)

	return nil
}
