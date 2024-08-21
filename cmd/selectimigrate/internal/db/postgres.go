package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/yourusername/selectimigrate/internal/config"
)

type PostgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase(cfg config.DBConfig) (*PostgresDatabase, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Credentials.Username, cfg.Credentials.Password, cfg.Database)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresDatabase{db: db}, nil
}

func (pdb *PostgresDatabase) Close() error {
	return pdb.db.Close()
}

// Implement other database operations here
