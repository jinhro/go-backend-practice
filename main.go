package main

import (
	"database/sql"
	"jinhro/go-backend-practice/api"
	db "jinhro/go-backend-practice/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable"
	serverAddress = "localhost:8088"
)

func main() {
	var err error

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}