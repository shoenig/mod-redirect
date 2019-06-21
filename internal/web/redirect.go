package web

import (
	"html/template"
	"net/http"

	"github.com/shoenig/loggy"
	"github.com/shoenig/mod-redirect/internal/store"
)

type redirectPage struct {
	Named  string
	Source string
}

type redirectEP struct {
	html  *template.Template
	store store.Storage
	log   loggy.Logger
}

func newRedirectEP(store store.Storage) http.Handler {
	return &redirectEP{
		store: store,
		log:   loggy.New("redirect"),
	}
}

func (h *redirectEP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Infof("serving a request!")
}

func (h *redirectEP) get(r *http.Request) {
	// need to open source this first ...
	// var ns, pkg string
	// if err := gorilla.Parse(r, gorilla.Schema {
	//
	// })
	//
	//
	// r.URL.Path
	// get the namespace and module from r.Path
	// h.store.Get()
}
