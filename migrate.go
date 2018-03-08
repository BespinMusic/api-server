package main

import (
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func migrate() {
	m, err := migrate.New("file://./migrations", "postgres://localhost:5432/database?user=postgres&password=test&dbname=bespin&sslmode=disable")
	m.Migrate(1)
}
