package migrator

import (
	"github.com/yourusername/selectimigrate/internal/config"
)

type Column struct {
	config config.Column
}

func NewColumn(cfg config.Column) *Column {
	return &Column{config: cfg}
}

func (c *Column) Migrate(sourceDB, targetDB *db.Database) error {
	// Implement column migration logic here
	return nil
}
