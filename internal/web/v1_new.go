package web

import (
	"encoding/json"
	"net/http"

	"gophers.dev/cmds/mod-redirect/internal/mods"
	"gophers.dev/cmds/mod-redirect/internal/store"

	"gophers.dev/pkgs/loggy"
)

type newEP struct {
	store store.Storage
	log   loggy.Logger
}

func newNewEP(store store.Storage) http.Handler {
	return &newEP{
		store: store,
		log:   loggy.New("new-ep"),
	}
}

func msg(err error) string {
	if err != nil {
		return err.Error()
	}
	return "ok"
}

func (h *newEP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// e.g. POST /v1/new -d <mods.Redirection>

	code, err := h.post(r)
	http.Error(w, msg(err), code)
}

func (h *newEP) post(r *http.Request) (int, error) {

	// extract the redirection from the request
	var redirection mods.Redirection
	if err := json.NewDecoder(r.Body).Decode(&redirection); err != nil {
		h.log.Errorf("unable to decode request: %v", err)
		return http.StatusBadRequest, err
	}

	// make sure all the fields were provided
	if err := mods.Valid(&redirection); err != nil {
		h.log.Errorf("redirection not valid: %v", err)
		return http.StatusBadRequest, err
	}

	// and now put it in the store
	if err := h.store.Set(&redirection); err != nil {
		h.log.Errorf("unable to save redirection: %v", err)
		return http.StatusInternalServerError, err
	}

	h.log.Infof("added new redirection: %s", redirection)

	return http.StatusOK, nil
}
