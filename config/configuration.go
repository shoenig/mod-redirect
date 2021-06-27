package config

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gophers.dev/cmds/mod-redirect/internal/mods"
)

var (
	ErrNoBindAddress = errors.New("no bind address")
	ErrPortRange     = errors.New("port not within range")
)

const (
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 10 * time.Second
)

type Configuration struct {
	Domain    string         `json:"domain"`
	WebServer WebServer      `json:"web_server"`
	Modules   mods.Redirects `json:"modules"`
}

type WebServer struct {
	BindAddress    string `json:"bind_address"`
	Port           int    `json:"port"`
	ReadTimeoutMS  int    `json:"read_timeout_ms"`
	WriteTimeoutMS int    `json:"write_timeout_ms"`
}

func (s WebServer) Address() string {
	return fmt.Sprintf("%s:%d", s.BindAddress, s.Port)
}

func (s WebServer) Server(mux http.Handler) (*http.Server, error) {
	if s.BindAddress == "" {
		return nil, ErrNoBindAddress
	}

	if s.Port <= 1024 {
		return nil, ErrPortRange
	}

	readTimeout := ms(s.ReadTimeoutMS)
	if readTimeout == 0 {
		readTimeout = defaultReadTimeout
	}

	writeTimeout := ms(s.WriteTimeoutMS)
	if writeTimeout == 0 {
		writeTimeout = defaultWriteTimeout
	}

	address := fmt.Sprintf("%s:%d", s.BindAddress, s.Port)

	server := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server, nil
}

func ms(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}
