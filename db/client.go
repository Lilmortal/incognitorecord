package db

import (
	"time"
)

// PostTableCreator is a generic interface for handling database related configurations.
type PostTableCreator interface {
	CreateTable(interface{}) error
}

// PostWriter handles the creation of a post.
type PostWriter interface {
	CreatePost(creationDate time.Time, title string, post string) error
}

// PostReader handles getting the post.
type PostReader interface {
	GetPost(title string) error
}

// PostDeleter handles the deletion of a post.
type PostDeleter interface {
	DeletePost(title string) error
}

// PostClient is a combination of all posts related interfaces
// TODO: The convention names for golang interfaces are that they ends with "er". e.g. Reader. Think of a good name
type PostClient interface {
	PostTableCreator
	PostWriter
	PostReader
	PostDeleter
}
