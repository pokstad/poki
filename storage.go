package poki

import "context"

type Storage interface {
	Create(context.Context, Post) (PostRev, error)
	Read(context.Context, string) (PostRev, error)
	Update(context.Context, Post, PostRev) (PostRev, error)
	Remove(context.Context, PostRev) error
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
