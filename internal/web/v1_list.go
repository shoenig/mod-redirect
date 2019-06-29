package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/shoenig/loggy"
	"github.com/shoenig/mod-redirect/internal/store"
)

type listEP struct {
	store store.Storage
	log   loggy.Logger
}

func newListEP(store store.Storage) http.Handler {
	return &listEP{
		store: store,
		log:   loggy.New("list-ep"),
	}
}

func (h *listEP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// e.g. GET /v1/list

	code, err := h.get(r)
	http.Error(w, msg(err), code)
}

func (h *listEP) get(r *http.Request) (int, error) {
	redirects, err := h.store.List()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	bs, err := json.Marshal(redirects)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, errors.New(string(bs))
}
