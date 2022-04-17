package main

import (
	"database/sql"
	"log"

	api "github.com/ardaatahan/simplebank/api"
	db "github.com/ardaatahan/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.RunServer(address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}