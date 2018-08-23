package db

import (
	"time"
)

// DatabaseClient is a generic interface that will be implemented by different databases, it will handle posts.
// TODO: The convention names for golang interfaces are that they ends with "er". e.g. Reader. Think of a good name
type DatabaseClient interface {
	CreateTable() error
}

type PostCreator interface {
	CreatePost(creationDate time.Time, title string, post string) error
}

// PostReader
