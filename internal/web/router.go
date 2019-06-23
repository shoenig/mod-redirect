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
	serveMux *http.ServeMux,
	storage store.Storage,
) {
	sub := mux.NewRouter()

	// namespace something like pkgs, cmds, src, etc.
	// or whatever your heart desires
	sub.Handle("/{namespace}/{module}", newRedirectEP(storage)).Methods(get)

	// implement this laterz
	// sub.Handle("/v1/set", newSetEP()).Methods(post)

	serveMux.Handle("/", sub)
}
