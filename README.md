# SelectiMigrate

SelectiMigrate is a flexible and powerful database migration tool designed for PostgreSQL. It allows for selective migration of tables and data, with customizable options for each migration task.

## Features

- Selective table and data migration
- Customizable migration options per table
- Parallel migration support
- Data transformation capabilities
- Pre and post-migration validations
- Detailed logging and error handling

## Installation

```bash
go get github.com/yourusername/selectimigrate
```

## Usage

1. Create a configuration file (see `examples/config.json` for a sample).
2. Run the migration:

```bash
selectimigrate --config path/to/your/config.json
```

## Configuration

The configuration file is in JSON format. Here's a basic structure:

```json
{
  "migrationName": "YourMigrationName",
  "source": {
    "type": "PostgreSQL",
    "host": "source-host",
    "port": 5432,
    "database": "source-db",
    "credentials": {
      "username": "source-user",
      "password": "source-password"
    }
  },
  "target": {
    "type": "PostgreSQL",
    "host": "target-host",
    "port": 5432,
    "database": "target-db",
    "credentials": {
      "username": "target-user",
      "password": "target-password"
    }
  },
  "tables": [
    {
      "sourceName": "source_table",
      "targetName": "target_table",
      "migrationOptions": {
        "migrateStructure": true,
        "migrateData": true
      },
      "columns": [
        {
          "source": "source_column",
          "target": "target_column",
          "type": "VARCHAR(255)"
        }
      ]
    }
  ]
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
