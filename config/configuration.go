package config

import (
	"net/http"
)

type Configuration struct {
	WebServer WebServer `json:"web_server"`
}

type WebServer struct {
	BindAddress    string `json:"bind_address"`
	Port           int    `json:"port"`
	ReadTimeoutMS  int    `json:"read_timeout_ms"`
	WriteTimeoutMS int    `json:"write_timeout_ms"`
}

func (s WebServer) Server(mux http.Handler) (*http.Server, error) {
	return nil, nil
}
