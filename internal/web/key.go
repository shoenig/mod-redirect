package web

import (
	"fmt"
	"net/http"
)

type SharedKey struct {
	header string
	keys   []string
}

func NewSharedKey(header string, keys []string) *SharedKey {

	if header == "" {
		panic("header is required")
	}

	if len(keys) == 0 {
		panic("one or more keys required")
	}

	return &SharedKey{
		header: header,
		keys:   keys,
	}
}

func (sk *SharedKey) CheckKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		givenKey := r.Header.Get(sk.header)

		for _, key := range sk.keys {
			if givenKey == key {
				next.ServeHTTP(w, r)
				return
			}
		}

		msg := fmt.Sprintf("header %s is wrong", sk.header)
		http.Error(w, msg, http.StatusBadRequest)
	})
}
