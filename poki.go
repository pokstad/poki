package poki

import (
	"context"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MetaData struct {
	Title string
}

type Post struct {
	Path string
	Raw  []byte
	Meta MetaData
}

type PostRev struct {
	Post
	RevisionID string
}

type Storage interface {
	Create(context.Context, Post) (PostRev, error)
	Read(context.Context, string) (PostRev, error)
	Update(context.Context, Post, PostRev) error
	Remove(context.Context, PostRev) error
}

type MemoryStorage struct {
	l sync.RWMutex
	m map[string]PostRev // revisions mapped by post path
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		m: map[string]PostRev{},
	}
}

type errCode int

const (
	ErrAlreadyExists errCode = 1 // Already exists
)

type Error struct {
	Code errCode
}

func (e Error) Error() string { return string(e.Code) }

func (ms *MemoryStorage) Create(ctx context.Context, p Post) (PostRev, error) {
	// check if it already exists
	ms.l.RLock()
	_, ok := ms.m[p.Path]
	if ok {
		// It already exists :(
		return PostRev{}, Error{Code: ErrAlreadyExists}
	}
	ms.l.RUnlock()

	// add it
	ms.l.Lock()
	defer ms.l.Unlock()
	rev := PostRev{
		Post:       p,
		RevisionID: uuid.NewV4().String(), // random ID
	}
	ms.m[p.Path] = rev

	return rev, nil
}

func (ms *MemoryStorage) Read(context.Context, string) (PostRev, error) {
	return PostRev{}, nil
}

func (ms *MemoryStorage) Update(context.Context, Post, PostRev) error {
	return nil
}

func (ms *MemoryStorage) Remove(context.Context, PostRev) error {
	return nil
}
