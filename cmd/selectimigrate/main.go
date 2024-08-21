package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yourusername/selectimigrate/internal/config"
	"github.com/yourusername/selectimigrate/internal/migrator"
)

func main() {
	// Define command-line flags
	configFile := flag.String("config", "", "Path to the configuration file")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	// Validate that config file is provided
	if *configFile == "" {
		fmt.Println("Error: Configuration file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up logging
	if *verbose {
		// TODO: Implement verbose logging
		fmt.Println("Verbose logging enabled")
	}

	// Initialize migrator
	m, err := migrator.NewMigrator(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize migrator: %v", err)
	}

	// Run migration
	err = m.Migrate()
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully")
}
