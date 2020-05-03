package web

import (
	"net/http"

	"gophers.dev/pkgs/loggy"
)

type healthEP struct {
	log loggy.Logger
}

func newHealthEP() http.Handler {
	return &healthEP{
		log: loggy.New("health"),
	}
}

func (h *healthEP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	http.Error(w, "ok", http.StatusOK)
}
