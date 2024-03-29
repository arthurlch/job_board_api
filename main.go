package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/arthurlch/job_board_api/api"
	db "github.com/arthurlch/job_board_api/db/sqlc"
)
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/jobb_dev?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
conn, err := sql.Open(dbDriver, dbSource)

if err != nil {
	log.Fatal("cannot connect to the db", err)
}

store := db.NewStore(conn)
server := api.NewServer(store)

err = server.Start(serverAddress)

if err != nil {
	log.Fatal("cannot start the server", err)
}
}