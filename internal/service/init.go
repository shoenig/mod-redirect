package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/connect"

	"gophers.dev/cmds/mod-redirect/config"
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
	address := config.Address()
	r.log.Tracef("setting up web service %s @ %s", config.Service(), address)

	domain := r.config.Domain
	if domain == "" {
		return errors.New("domain must be specified")
	}

	router := mux.NewRouter()
	web.Set(router, domain, r.storage)

	service := config.Service()
	consul := config.Consul()
	cs, err := connect.NewService(service, consul)
	if err != nil {
		return fmt.Errorf("unable to create consul service: %w", err)
	}

	go func(h http.Handler) {
		panic((&http.Server{
			Addr:              address,
			Handler:           router,
			TLSConfig:         cs.ServerTLSConfig(),
			ReadTimeout:       1 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,
			WriteTimeout:      1 * time.Second,
		}).ListenAndServeTLS("", ""))
	}(router)

	return nil
}
