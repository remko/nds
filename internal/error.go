package nds

import (
	"errors"
	"fmt"
)

type ndsError struct {
	kind    error
	wrapped error
}

func NewError(kind error, wrapped error) error {
	return &ndsError{kind: kind, wrapped: wrapped}
}

func (e *ndsError) Error() string {
	if e.kind == nil && e.wrapped == nil {
		return "db error"
	} else if e.kind == nil {
		return e.wrapped.Error()
	} else if e.wrapped == nil {
		return e.kind.Error()
	}
	return fmt.Sprintf("%s: %s", e.kind.Error(), e.wrapped.Error())
}

func (e *ndsError) Unwrap() error {
	return e.wrapped
}

func (e *ndsError) Is(err error) bool {
	return errors.Is(e.kind, err)
}
