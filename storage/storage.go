package storage

import (
	"context"

	"github.com/pokstad/poki"
)

type Storage interface {
	Create(context.Context, poki.Post) (poki.PostRev, error)
	Read(context.Context, string) (poki.PostRev, error)
	Update(context.Context, poki.Post, poki.PostRev) (poki.PostRev, error)
	Remove(context.Context, poki.PostRev) error
}

type errCode int

const (
	ErrAlreadyExists errCode = 1 // Already exists
	ErrNotFound      errCode = 2 // Not Found
	ErrConflict      errCode = 3 // Conflict
	ErrInvalidParams errCode = 4 // Invalid Parameters
)

type Error struct {
	Code errCode
}

func (e Error) Error() string { return string(e.Code) }
