package cache

import (
	"context"
	"io"
)

type FetchFunc func() (interface{}, error)

type CacheInterface interface {
	io.Closer

	Fetch(ctx context.Context, key string, value interface{}, fetchFunc FetchFunc) error

	Get(ctx context.Context, key string, value interface{}) error

	Set(ctx context.Context, key string, value interface{}) error

	Exists(ctx context.Context, key string) (bool, error)

	Delete(ctx context.Context, key string) error
}
