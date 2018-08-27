package db

import (
	"time"
)

// DB is a generic interface that will be implemented by different databases, it will handle posts.
// TODO: The convention names for golang interfaces are that they ends with "er". e.g. Reader. Think of a good name
type DB interface {
	CreateTable() error
}

type PostWriter interface {
	CreatePost(creationDate time.Time, title string, post string) error
}

type PostReader interface {
	GetPost(title string) error
}

type PostDB interface {
	CreateTable() error
	CreatePost(creationDate time.Time, title string, post string) error
	GetPost(title string) error
}
