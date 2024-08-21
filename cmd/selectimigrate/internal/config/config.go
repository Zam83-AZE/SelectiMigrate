package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	MigrationName string    `json:"migrationName"`
	Description   string    `json:"description"`
	Version       string    `json:"version"`
	CreatedAt     time.Time `json:"createdAt"`
	Author        string    `json:"author"`
	MigrationType string    `json:"migrationType"`
	Source        DBConfig  `json:"source"`
	Target        DBConfig  `json:"target"`
	Tables        []Table   `json:"tables"`
	DataTransformations []DataTransformation `json:"dataTransformations"`
	Validations   []Validation `json:"validations"`
	ErrorHandling ErrorHandling `json:"errorHandling"`
	Logging       Logging       `json:"logging"`
	PostMigrationScripts []string `json:"postMigrationScripts"`
}

type DBConfig struct {
	Type        string `json:"type"`
	Version     string `json:"version"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Database    string `json:"database"`
	Schema      string `json:"schema"`
	Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
}

type Table struct {
	SourceName      string         `json:"sourceName"`
	TargetName      string         `json:"targetName"`
	PrimaryKey      string         `json:"primaryKey"`
	DependsOn       []string       `json:"dependsOn"`
	MigrationOptions MigrationOptions `json:"migrationOptions"`
	Columns         []Column       `json:"columns"`
	Transformations []Transformation `json:"transformations"`
	IndexesToCreate []Index        `json:"indexesToCreate"`
}

type MigrationOptions struct {
	MigrateStructure    bool   `json:"migrateStructure"`
	MigrateData         bool   `json:"migrateData"`
	TruncateBeforeInsert bool  `json:"truncateBeforeInsert"`
	DataFilter           string `json:"dataFilter"`
	MigrationOperation   string `json:"migrationOperation"`
	ParallelMigration    bool   `json:"parallelMigration"`
	ParallelizationFactor int   `json:"parallelizationFactor"`
}

type Column struct {
	Source       string `json:"source"`
	Target       string `json:"target"`
	Type         string `json:"type"`
	AutoIncrement bool  `json:"autoIncrement"`
	Unique       bool   `json:"unique"`
	ForeignKey   *ForeignKey `json:"foreignKey"`
}

type ForeignKey struct {
	Table  string `json:"table"`
	Column string `json:"column"`
}

type Transformation struct {
	Column string `json:"column"`
	Action string `json:"action"`
}

type Index struct {
	Name    string   `json:"name"`
	Columns []string `json:"columns"`
}

type DataTransformation struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	TargetTable string `json:"targetTable"`
	SQL         string `json:"sql"`
}

type Validation struct {
	Type         string `json:"type"`
	Description  string `json:"description"`
	SourceQuery  string `json:"sourceQuery"`
	TargetQuery  string `json:"targetQuery"`
	ExpectedResult interface{} `json:"expectedResult"`
}

type ErrorHandling struct {
	OnError      string `json:"onError"`
	RetryAttempts int    `json:"retryAttempts"`
	RetryDelay   int    `json:"retryDelay"`
}

type Logging struct {
	Level       string `json:"level"`
	Destination string `json:"destination"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	err = validateConfig(&config)
	if err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}

	return &config, nil
}

func validateConfig(config *Config) error {
	if config.MigrationName == "" {
		return fmt.Errorf("migration name is required")
	}
	if config.Source.Type == "" || config.Target.Type == "" {
		return fmt.Errorf("source and target database types are required")
	}
	if len(config.Tables) == 0 {
		return fmt.Errorf("at least one table configuration is required")
	}
	// Add more validation as needed
	return nil
}
