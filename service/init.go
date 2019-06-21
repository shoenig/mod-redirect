package service

import (
	"net/http"

	"github.com/shoenig/mod-redirect/internal/web"
)

type initer func(*Proxy) error

func initWebServer(p *Proxy) error {
	smux := http.NewServeMux()
	web.Set(smux)
	return nil
}
