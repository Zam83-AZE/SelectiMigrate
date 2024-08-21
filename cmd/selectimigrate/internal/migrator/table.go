package migrator

import (
	"fmt"
	"github.com/yourusername/selectimigrate/internal/config"
	"github.com/yourusername/selectimigrate/internal/db"
)

type Migrator struct {
	config *config.Config
	sourceDB *db.Database
	targetDB *db.Database
}

func NewMigrator(cfg *config.Config) (*Migrator, error) {
	sourceDB, err := db.NewDatabase(cfg.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to source database: %w", err)
	}

	targetDB, err := db.NewDatabase(cfg.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to target database: %w", err)
	}

	return &Migrator{
		config: cfg,
		sourceDB: sourceDB,
		targetDB: targetDB,
	}, nil
}

func (m *Migrator) Migrate() error {
	for _, table := range m.config.Tables {
		err := m.migrateTable(table)
		if err != nil {
			return fmt.Errorf("failed to migrate table %s: %w", table.SourceName, err)
		}
	}

	return nil
}

func (m *Migrator) migrateTable(table config.Table) error {
	// Implement table migration logic here
	return nil
}
