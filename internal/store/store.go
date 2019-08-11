package store

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"gophers.dev/cmds/mod-redirect/internal/mods"
)

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -g -i Storage -s _mock.go

type Storage interface {
	Set(*mods.Redirection) error
	Get(string, string) (*mods.Redirection, error)
	List() ([]*mods.Redirection, error)
}

type boltDB struct {
	filepath string
	db       *bolt.DB
}

func NewBoltDB(filepath string) (Storage, error) {
	db, err := bolt.Open(filepath, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "could not open storage")
	}

	if err := setup(db); err != nil {
		return nil, err
	}

	return &boltDB{
		filepath: filepath,
		db:       db,
	}, nil
}

func setup(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(redirectBktLbl)
		return err
	})
}

var (
	redirectBktLbl = []byte("redirects")
)

func (b *boltDB) Set(r *mods.Redirection) error {
	key := r.Key()
	value := r.Bytes()

	return b.db.Update(func(tx *bolt.Tx) error {
		redirectBkt := tx.Bucket(redirectBktLbl)
		return redirectBkt.Put(key, value)
	})
}

var (
	ErrModuleDoesNotExist = errors.New("module does not exist")
)

func (b *boltDB) Get(kind, module string) (*mods.Redirection, error) {
	key := mods.Key(kind, module)

	var redirect mods.Redirection
	var content []byte

	if err := b.db.View(func(tx *bolt.Tx) error {
		redirectBkt := tx.Bucket(redirectBktLbl)
		bs := redirectBkt.Get(key) // must copy inside tx
		content = make([]byte, len(bs))
		copy(content, bs)
		if bs == nil {
			return ErrModuleDoesNotExist
		}
		return nil
	}); err != nil {
		return nil, err
	}

	err := json.Unmarshal(content, &redirect)
	return &redirect, err
}

func (b *boltDB) List() ([]*mods.Redirection, error) {
	var redirects []*mods.Redirection

	if err := b.db.View(func(tx *bolt.Tx) error {
		redirectBkt := tx.Bucket(redirectBktLbl)
		return redirectBkt.ForEach(func(_, v []byte) error {
			var redirect mods.Redirection
			if err := json.Unmarshal(v, &redirect); err != nil {
				return err
			}
			redirects = append(redirects, &redirect)
			return nil
		})
	}); err != nil {
		return nil, err
	}

	return redirects, nil
}
