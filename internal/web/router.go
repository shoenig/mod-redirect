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
	sub *mux.Router,
	storage store.Storage,
) {

	// needs to be protected with a key or something
	sub.Handle("/v1/set", newNewEP(storage)).Methods(post)
	// sub.Handle("/v1/list", newListEP(storage)).Methods(get)

	// namespace something like pkgs, cmds, src, etc.
	// the module could be anything after that (word characters and slash)
	sub.Handle(`/{namespace}/{module:[a-zA-Z0-9/_-]+}`, newRedirectEP(storage)).Methods(get)
}
