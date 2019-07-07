package web

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"

	"github.com/shoenig/extractors/urlpath"
	"github.com/shoenig/loggy"
	"github.com/shoenig/mod-redirect/internal/mods"
	"github.com/shoenig/mod-redirect/internal/store"
	"github.com/shoenig/mod-redirect/static"
)

type redirectPage struct {
	Domain string
	Kind   string
	Named  string
	VCS    string
	Source string
}

type redirectEP struct {
	html   *template.Template
	domain string
	store  store.Storage
	log    loggy.Logger
}

func newRedirectEP(domain string, store store.Storage) http.Handler {
	html := static.MustParseTemplates(
		"static/html/goget.html",
	)

	return &redirectEP{
		html:   html,
		domain: domain,
		store:  store,
		log:    loggy.New("redirect"),
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
		h.log.Errorf("unable to extract module %s, %v", module, err)
		return http.StatusBadRequest, err
	}

	redirection, err := h.store.Get(namespace, module)
	if err != nil {
		h.log.Errorf("unable to fetch redirection: %v", err)
		return http.StatusInternalServerError, err
	}

	h.log.Infof("redirect %s/%s -> %s", namespace, module, redirection)

	response, err := h.render(redirection)
	if err != nil {
		h.log.Errorf("unable to render response: %v", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, errors.New(response)
}

func (h *redirectEP) render(redirection *mods.Redirection) (string, error) {
	var content bytes.Buffer
	if err := h.html.Execute(&content, redirectPage{
		Domain: h.domain,
		Kind:   redirection.Kind,
		Named:  redirection.Namespace,
		VCS:    redirection.VCS,
		Source: redirection.Destination,
	}); err != nil {
		return "", err
	}
	return content.String(), nil
}
