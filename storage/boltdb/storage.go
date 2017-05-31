package boltdb

import (
	"context"

	"crypto/sha256"

	"encoding/base64"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/pokstad/poki"
)

// RabbitStorage is a storage engine based on boltdb
type RabbitStorage struct {
	*bolt.DB
}

// NewDatabase opens an existing database or creates a new one
func NewDatabase(filepath string) (*RabbitStorage, error) {
	db, err := bolt.Open(filepath, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &RabbitStorage{
		DB: db,
	}, nil
}

func (rs *RabbitStorage) Create(ctx context.Context, p poki.Post) (poki.PostRev, error) {
	var (
		hash []byte
	)

	err := rs.DB.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucket([]byte(p.Path))
		if err != nil && err != bolt.ErrBucketExists {
			return errors.Wrap(err, "could not create bucket")
		}

		// Store post with sha256 of post content as key
		hash := sha256.Sum256(p.Raw)

		if err := b.Put(hash[:], p.Raw); err != nil {
			return errors.Wrap(err, "could not put post into bucket")
		}

		return nil
	})

	return poki.PostRev{
		Post:       p,
		RevisionID: base64.StdEncoding.EncodeToString(hash),
	}, err
}

func (rs *RabbitStorage) Read(context.Context, string) (poki.PostRev, error) {
	return poki.PostRev{}, nil
}

func (rs *RabbitStorage) Update(context.Context, poki.Post, poki.PostRev) (poki.PostRev, error) {
	return poki.PostRev{}, nil
}

func (rs *RabbitStorage) Remove(context.Context, poki.PostRev) error { return nil }
