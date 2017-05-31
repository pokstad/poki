package memory

import (
	"context"
	"sync"

	"github.com/pokstad/poki"
	"github.com/pokstad/poki/storage"
	uuid "github.com/satori/go.uuid"
)

type MemoryStorage struct {
	l sync.RWMutex
	m map[string]poki.PostRev // revisions mapped by post path
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		m: map[string]poki.PostRev{},
	}
}

func (ms *MemoryStorage) Create(_ context.Context, p poki.Post) (poki.PostRev, error) {
	// check if it already exists
	ms.l.RLock()
	_, ok := ms.m[p.Path]
	ms.l.RUnlock()

	if ok {
		// It already exists :(
		return poki.PostRev{}, storage.Error{Code: storage.ErrAlreadyExists}
	}

	// add it
	rev := poki.PostRev{
		Post:       p,
		RevisionID: uuid.NewV4().String(), // random ID
	}

	ms.l.Lock()
	ms.m[p.Path] = rev
	ms.l.Unlock()

	return rev, nil
}

func (ms *MemoryStorage) Read(_ context.Context, path string) (poki.PostRev, error) {
	ms.l.RLock()
	post, ok := ms.m[path]
	ms.l.RUnlock()

	if !ok {
		return poki.PostRev{}, storage.Error{Code: storage.ErrNotFound}
	}

	return post, nil
}

func (ms *MemoryStorage) Update(_ context.Context, newPost poki.Post, prev poki.PostRev) (poki.PostRev, error) {
	// validate params
	if newPost.Path != prev.Path {
		// invalid params
		return poki.PostRev{}, storage.Error{Code: storage.ErrInvalidParams}
	}

	// fetch current revision
	ms.l.RLock()
	currentRev, ok := ms.m[prev.Post.Path]
	ms.l.RUnlock()

	switch {

	// exists and revs mismatch
	case ok && currentRev.RevisionID != prev.RevisionID:
		return poki.PostRev{}, storage.Error{
			Code: storage.ErrConflict,
		}

	// exists and revs match
	case ok && currentRev.RevisionID == prev.RevisionID:
		newRev := poki.PostRev{
			Post:       newPost,
			RevisionID: uuid.NewV4().String(), // random ID
		}

		ms.l.Lock()
		ms.m[prev.Path] = newRev
		ms.l.Unlock()

		return poki.PostRev{}, nil

	// doesn't exist
	case !ok:
		newRev := poki.PostRev{
			Post:       newPost,
			RevisionID: uuid.NewV4().String(), // random ID
		}

		ms.l.Lock()
		ms.m[prev.Path] = newRev
		ms.l.Unlock()

		return newRev, nil
	}

	return poki.PostRev{}, nil
}

func (ms *MemoryStorage) Remove(_ context.Context, _ poki.PostRev) error {
	return nil
}
