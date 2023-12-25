package cache

import (
	"errors"
	"strings"
)

var (
	ErrCacheMiss    = errors.New("key is missing")
	ErrInvalidKey   = errors.New("key is invalid")
	ErrInvalidValue = errors.New("value type is invalid")
)

func wrapError(err error) error {
	if err == nil {
		return nil
	}

	switch err {
	case ErrCacheMiss:
		return ErrCacheMiss
	}

	msg := err.Error()

	switch {
	case strings.Contains(msg, "unknow compression method"):
		return ErrInvalidKey
	}
	return err
}
