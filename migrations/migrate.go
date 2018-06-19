package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func formatURL(host, port, user, password, dbname string) string {
	connectionURL :=
		fmt.Sprintf("postgres://%s:%s/%s?user=%s&password=%s&sslmode=disable&dbname=%s", host, port, dbname, user, password, dbname)

	return connectionURL
}

func main() {

	connectionURL := formatURL(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	log.Print(connectionURL)
	m, err := migrate.New("file://./migrations", connectionURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	m.Migrate(1)
	return
}
