package storage

import "context"

// Storage is an interface for all backup provider
type Storage interface {
	Name() string

	Backup(ctx context.Context, source string, targetSubPath string) error

	Validate() error

	Close() error
}

type Config struct {
	Type string
	Path string
}
