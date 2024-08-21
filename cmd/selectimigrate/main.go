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
	configFile := flag.String("config", "", "Path to the configuration file")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("Error: Configuration file is required")
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if *verbose {
		// TODO: Implement verbose logging
		fmt.Println("Verbose logging enabled")
	}

	m, err := migrator.NewMigrator(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize migrator: %v", err)
	}

	err = m.Migrate()
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully")
}
