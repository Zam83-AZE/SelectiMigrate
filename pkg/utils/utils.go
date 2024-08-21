package utils

import (
	"fmt"
	"strings"
)

func FormatColumnName(name string) string {
	return strings.ToLower(name)
}

func GenerateCreateTableSQL(tableName string, columns []string) string {
	return fmt.Sprintf("CREATE TABLE %s (%s)", tableName, strings.Join(columns, ", "))
}

// Add more utility functions as needed
