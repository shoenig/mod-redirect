package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shoenig/mod-redirect/internal/store"
)

const (
	get  = http.MethodGet
	post = http.MethodPost
)

func Set(
	serveMux *mux.Router,
	storage store.Storage,
) {
	sub := mux.NewRouter()

	// needs to be protected with a key or something
	sub.Handle("/v1/set", newNewEP(storage)).Methods(post)

	// namespace something like pkgs, cmds, src, etc.
	// or whatever your heart desires
	sub.Handle("/{namespace}/{module}", newRedirectEP(storage)).Methods(get)

	serveMux.Handle("/", sub)
}
