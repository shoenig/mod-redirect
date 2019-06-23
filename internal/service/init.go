package service

import (
	"github.com/pkg/errors"
	"github.com/shoenig/mod-redirect/internal/store"
)

type initer func(*Redirect) error

func initStore(r *Redirect) error {
	fPath := r.config.Storage.BoltDB.Filepath
	r.log.Tracef("setting up boltdb with path %s", fPath)
	boltDB, err := store.NewBoltDB(fPath)
	if err != nil {
		return errors.Wrap(err, "unable to create boltdb store")
	}
	r.storage = boltDB
	return nil
}

func initWeb(r *Redirect) error {
	r.log.Tracef("setting up web server")

	// hello

	return nil
}
