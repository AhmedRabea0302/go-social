package main

import (
	"log"

	"github.com/AhmedRabea0302/go-social/internal/db"
	"github.com/AhmedRabea0302/go-social/internal/env"
	"github.com/AhmedRabea0302/go-social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:0123456@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()
	log.Println("database connetion pool initialized")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	// Add Registered Routes
	mux := app.mount()

	log.Fatal(app.run(mux))
}
