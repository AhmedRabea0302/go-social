package main

import (
	"log"

	"github.com/AhmedRabea0302/go-social/internal/env"
	"github.com/AhmedRabea0302/go-social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	// Add Registered Routes
	mux := app.mount()

	log.Fatal(app.run(mux))
}
