package poki

import (
	"context"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemoryStorage struct {
	l sync.RWMutex
	m map[string]PostRev // revisions mapped by post path
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		m: map[string]PostRev{},
	}
}

func (ms *MemoryStorage) Create(_ context.Context, p Post) (PostRev, error) {
	// check if it already exists
	ms.l.RLock()
	_, ok := ms.m[p.Path]
	ms.l.RUnlock()

	if ok {
		// It already exists :(
		return PostRev{}, Error{Code: ErrAlreadyExists}
	}

	// add it
	rev := PostRev{
		Post:       p,
		RevisionID: uuid.NewV4().String(), // random ID
	}

	ms.l.Lock()
	ms.m[p.Path] = rev
	ms.l.Unlock()

	return rev, nil
}

func (ms *MemoryStorage) Read(_ context.Context, path string) (PostRev, error) {
	ms.l.RLock()
	post, ok := ms.m[path]
	ms.l.RUnlock()

	if !ok {
		return PostRev{}, Error{Code: ErrNotFound}
	}

	return post, nil
}

func (ms *MemoryStorage) Update(_ context.Context, newPost Post, prev PostRev) (PostRev, error) {
	// validate params
	if newPost.Path != prev.Path {
		// invalid params
		return PostRev{}, Error{Code: ErrInvalidParams}
	}

	// fetch current revision
	ms.l.RLock()
	currentRev, ok := ms.m[prev.Post.Path]
	ms.l.RUnlock()

	switch {

	// exists and revs mismatch
	case ok && currentRev.RevisionID != prev.RevisionID:
		return PostRev{}, Error{
			Code: ErrConflict,
		}

	// exists and revs match
	case ok && currentRev.RevisionID == prev.RevisionID:
		newRev := PostRev{
			Post:       newPost,
			RevisionID: uuid.NewV4().String(), // random ID
		}

		ms.l.Lock()
		ms.m[prev.Path] = newRev
		ms.l.Unlock()

		return PostRev{}, nil

	// doesn't exist
	case !ok:
		newRev := PostRev{
			Post:       newPost,
			RevisionID: uuid.NewV4().String(), // random ID
		}

		ms.l.Lock()
		ms.m[prev.Path] = newRev
		ms.l.Unlock()

		return newRev, nil
	}

	return PostRev{}, nil
}

func (ms *MemoryStorage) Remove(context.Context, PostRev) error {
	return nil
}
