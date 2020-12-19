package database

import (
	"errors"
)

// ErrConnectionNotEstablished is returned when connection is not established but is needed
var ErrConnectionNotEstablished = errors.New("Connection to database not established")

// DB defines database operations
type DB interface {
	Connect() error
	MigrateUp() error
	MigrateDown() error
}

// Config defines database configuration
type Config struct {
	URL            string
	MigrationFiles string
}
