package web

import (
	"net/http"

	"github.com/gorilla/mux"

	"gophers.dev/cmds/mod-redirect/internal/store"
)

const (
	get = http.MethodGet
)

func Set(
	router *mux.Router,
	domain string,
	storage store.Storage,
) {

	// health endpoint
	router.Handle("/health", newHealthEP()).Methods(get)

	// api endpoints
	router.Handle("/v1/list", newListEP(storage)).Methods(get)

	// namespace something like pkgs, cmds, src, etc.
	// the module could be anything after that (word characters and slash)
	router.Handle(`/{namespace}/{module:[a-zA-Z0-9/_-]+}`, newRedirectEP(domain, storage)).Methods(get)
}

func msg(err error) string {
	if err != nil {
		return err.Error()
	}
	return "ok"
}
