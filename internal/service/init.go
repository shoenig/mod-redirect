package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/pkg/errors"

	"github.com/shoenig/mod-redirect/internal/store"
	"github.com/shoenig/mod-redirect/internal/web"
)

type initer func(*Redirect) error

func initStore(r *Redirect) error {
	fPath := r.config.Storage.BoltDB.Filepath
	r.log.Tracef("setting up boltdb with path %s", fPath)
	boltDB, err := store.NewBoltDB(fPath)
	if err != nil {
		return errors.Wrap(err, "unable to create boltdb store")
	}
	r.storage = boltDB
	return nil
}

func initWeb(r *Redirect) error {
	r.log.Tracef("setting up web server @ %s", r.config.WebServer.Address())

	header := r.config.Authentication.Header
	keys := r.config.Authentication.Keys
	sk := web.NewSharedKey(header, keys)

	router := mux.NewRouter()
	web.Set(router, r.storage, sk)

	server, err := r.config.WebServer.Server(router)
	if err != nil {
		return errors.Wrap(err, "could not create web server")
	}

	go func(h http.Handler) {
		err := server.ListenAndServe()
		r.log.Errorf("server stopped serving: %v", err)
		os.Exit(1)
	}(router)

	return nil
}
