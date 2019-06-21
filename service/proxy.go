package service

import (
	"github.com/shoenig/loggy"
	"github.com/shoenig/mod-redirect/config"
)

type Proxy struct {
	config config.Configuration
	log    loggy.Logger
}

func NewProxy(config config.Configuration) *Proxy {
	return &Proxy{
		config: config,
		log:    loggy.New("proxy"),
	}
}
