package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/remko/nds"
)

type DB struct {
	client *datastore.Client
}

type query struct {
	query *datastore.Query
}

func (q query) Start(c nds.Cursor) nds.Query {
	return query{q.query.Start(c)}
}

func (q query) Limit(limit int) nds.Query {
	return query{q.query.Limit(limit)}
}

func (q query) Filter(filterStr string, value interface{}) nds.Query {
	return query{q.query.Filter(filterStr, value)}
}

func NewDB(ctx context.Context, project string) (*DB, error) {
	client, err := datastore.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}
	return &DB{
		client: client,
	}, nil
}

func (db *DB) Close() {
	db.client.Close()
}

func (db *DB) NewQuery(kind string) nds.Query {
	return query{datastore.NewQuery(kind)}
}

func (db *DB) Get(ctx context.Context, key *nds.Key, dst interface{}) (err error) {
	return dsError(db.client.Get(ctx, key, dst))
}
func (db *DB) Put(ctx context.Context, key *nds.Key, src interface{}) (*nds.Key, error) {
	k, err := db.client.Put(ctx, key, src)
	return k, dsError(err)
}

func (db *DB) Run(ctx context.Context, q nds.Query) nds.Iterator {
	return db.client.Run(ctx, q.(query).query)
}
