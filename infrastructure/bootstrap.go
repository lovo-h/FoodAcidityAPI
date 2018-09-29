package infrastructure

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewCDBHandler() *HandlerCockroach {
	connStr := "postgresql://root@cockroach:26257/usdafooddb?sslmode=disable"
	db, dbErr := sql.Open("postgres", connStr)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	handlerCockroach := new(HandlerCockroach)
	handlerCockroach.DB = db

	return handlerCockroach
}