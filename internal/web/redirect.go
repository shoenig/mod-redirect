package web

import (
	"html/template"
	"net/http"

	"github.com/shoenig/extractors/urlpath"
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
	code, err := h.get(r)
	http.Error(w, msg(err), code)
}

func (h *redirectEP) get(r *http.Request) (int, error) {
	// e.g. GET example.com/pkgs/foo/bar

	var (
		namespace string
		module    string
	)

	if err := urlpath.Parse(r, urlpath.Schema{
		"namespace": urlpath.String(&namespace),
		"module":    urlpath.String(&module),
	}); err != nil {
		h.log.Errorf("unable to extract module: %v", err)
		return http.StatusBadRequest, err
	}

	redirection, err := h.store.Get(namespace, module)
	if err != nil {
		h.log.Errorf("unable to fetch redirection: %v", err)
		return http.StatusInternalServerError, err
	}

	h.log.Infof("redirect %s/%s -> %s", namespace, module, redirection)

	// generate redirect html

	return http.StatusOK, nil
}
