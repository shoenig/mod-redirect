package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	get  = http.MethodGet
	post = http.MethodPost
)

func Set(smux *http.ServeMux) {
	sub := mux.NewRouter()

	// namespace something like pkgs, cmds, src, etc.
	// or whatever your heart desires
	sub.Handle("/{namespace}/{module}", newRedirectEP()).Methods(get)

	// implement this laterz
	// sub.Handle("/v1/set", newSetEP()).Methods(post)

	smux.Handle("/", sub)
}
