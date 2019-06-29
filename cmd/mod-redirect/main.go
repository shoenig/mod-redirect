package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"github.com/shoenig/mod-redirect/config"
	"github.com/shoenig/mod-redirect/internal/service"
)

// generate HTML static file(s)
//go:generate go run github.com/shoenig/petrify/v4/cmd/petrify -prefix ../../ -o ../../static/generated.go -pkg static ../../static/...

func main() {
	flagSet := flag.NewFlagSet("configuration-flags", flag.ExitOnError)
	cfg, err := parse(flagSet, os.Args[1:])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not start: %v\n", err)
		os.Exit(1)
	}

	redirect := service.NewRedirect(cfg)
	redirect.Start()
}

func parse(flags *flag.FlagSet, args []string) (config.Configuration, error) {
	var (
		configFile    string
		configuration config.Configuration
	)

	flags.StringVar(&configFile, "config", "", "the configuration JSON file to read from")

	if err := flags.Parse(args); err != nil {
		return configuration, errors.Wrap(err, "unable to parse args")
	}

	if configFile == "" {
		return configuration, errors.New("--config [filename] is required")
	}

	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		return configuration, errors.Wrapf(err, "unable to read config file %q", configFile)
	}

	if err := json.Unmarshal(bs, &configuration); err != nil {
		return configuration, errors.Wrap(err, "unable to unmarshal config file")
	}

	return configuration, nil
}
