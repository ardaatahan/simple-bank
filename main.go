package main

import (
	"database/sql"
	"log"

	api "github.com/ardaatahan/simplebank/api"
	db "github.com/ardaatahan/simplebank/db/sqlc"
	"github.com/ardaatahan/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load environment variables: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.RunServer(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
