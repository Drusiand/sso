package migrator

import (
	"flag"

	"github.com/pressly/goose/v3"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if storagePath == ""{
		panic("storage-path is required")
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	
}
