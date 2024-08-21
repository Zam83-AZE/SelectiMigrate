package db

import (
	"fmt"
	"github.com/yourusername/selectimigrate/internal/config"
)

type Database interface {
	Close() error
	// Add other database operations here
}

func NewDatabase(cfg config.DBConfig) (Database, error) {
	switch cfg.Type {
	case "PostgreSQL":
		return NewPostgresDatabase(cfg)
	// Add cases for other database types here
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
}
