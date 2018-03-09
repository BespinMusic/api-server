package main

import (
	"log"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	m, err := migrate.New("file://./", "postgres://localhost:9145/bespin?user=postgres&password=test&dbname=bespin&sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
	m.Migrate(1)

}
