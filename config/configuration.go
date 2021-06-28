package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
	"gophers.dev/cmds/mod-redirect/internal/mods"
)

type Configuration struct {
	Domain  string         `json:"domain"`
	Modules mods.Redirects `json:"modules"`
}

// todo: helper methods

func mustGetInt(name string) int {
	if s := os.Getenv(name); s != "" {
		p, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(name + " must be a number")
		}
		return p
	}
	log.Fatal(name + " must be set")
	return -1
}

func getStringOr(name, alt string) string {
	if s := os.Getenv(name); s != "" {
		return s
	}
	return alt
}

func Address() string {
	port := mustGetInt("PORT")
	bind := getStringOr("BIND", "0.0.0.0")
	return fmt.Sprintf("%s:%d", bind, port)
}

func Service() string {
	return getStringOr("SERVICE", "mod-redirect")
}

func Consul() *api.Client {
	logEnvironment("CONSUL_HTTP_ADDR")
	logEnvironment("CONSUL_NAMESPACE")
	logEnvironment("CONSUL_CACERT")
	logEnvironment("CONSUL_CLIENT_CERT")
	logEnvironment("CONSUL_CLIENT_KEY")
	logEnvironment("CONSUL_HTTP_SSL")
	logEnvironment("CONSUL_HTTP_SSL_VERIFY")
	logEnvironment("CONSUL_TLS_SERVER_NAME")
	logEnvironment("CONSUL_GRPC_ADDR")
	logEnvironment("CONSUL_HTTP_TOKEN_FILE")
	consulConfig := api.DefaultConfig()
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal("failed to make consul client:", err)
	}
	return consulClient
}

func logEnvironment(name string) {
	value := os.Getenv(name)
	if value == "" {
		value = "<unset>"
	}
	log.Printf("environment %s = %s", name, value)
}

func ms(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}
