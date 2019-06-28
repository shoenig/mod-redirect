package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shoenig/loggy"
	"github.com/shoenig/mod-redirect/internal/mods"
	"github.com/shoenig/mod-redirect/internal/store"
)

type newEP struct {
	store store.Storage
	log   loggy.Logger
}

func newNewEP(store store.Storage) http.Handler {
	return &newEP{
		store: store,
		log:   loggy.New("new-redirct"),
	}
}

func msg(err error) string {
	if err != nil {
		return err.Error()
	}
	return "ok"
}

func (h *newEP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Infof("hi")
	// e.g. POST /v1/new -d <mods.Redirection>

	// todo: need some semblance of security

	code, err := h.post(r)
	http.Error(w, msg(err), code)
}

func (h *newEP) post(r *http.Request) (int, error) {
	// hello world

	var redirection mods.Redirection
	if err := json.NewDecoder(r.Body).Decode(&redirection); err != nil {
		return http.StatusBadRequest, err
	}

	fmt.Println("redirection:", redirection)

	return http.StatusOK, nil
}
