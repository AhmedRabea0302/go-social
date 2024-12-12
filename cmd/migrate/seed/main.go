package main

import (
	"log"

	"github.com/AhmedRabea0302/go-social/internal/db"
	"github.com/AhmedRabea0302/go-social/internal/env"
	"github.com/AhmedRabea0302/go-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://postgres:0123456@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 30, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}
