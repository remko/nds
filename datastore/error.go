package datastore

import (
	"errors"

	"cloud.google.com/go/datastore"
	"github.com/remko/nds"
	ndsi "github.com/remko/nds/internal"
)

func dsError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, datastore.ErrNoSuchEntity) {
		return ndsi.NewError(nds.ErrNotFound, err)
	}
	if _, ok := err.(*datastore.ErrFieldMismatch); ok {
		return ndsi.NewError(nds.ErrFieldMismatch, err)
	}
	return ndsi.NewError(nil, err)
}
