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
	router *mux.Router,
	domain string,
	storage store.Storage,
	checker *SharedKey,
) {

	router.Handle("/v1/set", setter(storage, checker))
	router.Handle("/v1/list", newListEP(storage)).Methods(get)

	// namespace something like pkgs, cmds, src, etc.
	// the module could be anything after that (word characters and slash)
	router.Handle(`/{namespace}/{module:[a-zA-Z0-9/_-]+}`, newRedirectEP(domain, storage)).Methods(get)
}

func setter(storage store.Storage, checker *SharedKey) http.Handler {
	sub := mux.NewRouter()
	sub.Use(checker.CheckKey)
	sub.Handle("/v1/set", newNewEP(storage)).Methods(post)
	return sub
}
