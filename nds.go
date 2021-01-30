// Package nds provides Google Cloud Datastore utility clients
package nds

import (
	"context"
	"errors"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

var ErrInvalidID = errors.New("invalid id")
var ErrInvalidKey = errors.New("invalid key")
var ErrNotFound = errors.New("not found")
var ErrFieldMismatch = errors.New("field mismatch")

// Primitive types directly taken from `cloud.google.com/go/datastore`
type Key = datastore.Key
type Cursor = datastore.Cursor

var IncompleteKey = datastore.IncompleteKey
var IteratorDone = iterator.Done

// Abstract interfaces for datastore types
type Iterator interface {
	Cursor() (c Cursor, err error)
	Next(dst interface{}) (k *Key, err error)
}

type Query interface {
	Filter(filterStr string, value interface{}) Query
	Start(c Cursor) Query
	Limit(limit int) Query
}

// Driver implementation.
// This implements the global functions of the `cloud.google.com/go/datastore` package.
type DB interface {
	NewQuery(kind string) Query

	Get(ctx context.Context, key *Key, dst interface{}) (err error)
	Put(ctx context.Context, key *Key, src interface{}) (*Key, error)
	Run(ctx context.Context, q Query) Iterator
}
