package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	dbname := os.Args[1]
	log.Print(dbname)

	connectionURL :=
		fmt.Sprintf("postgres://localhost:9145/%s?user=postgres&password=test&sslmode=disable&dbname=%s", dbname, dbname)

	log.Print(connectionURL)
	m, err := migrate.New("file://./migrations/", connectionURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	m.Migrate(1)
	return
}
