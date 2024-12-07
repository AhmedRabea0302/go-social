package main

import (
	"log"

	"github.com/AhmedRabea0302/go-social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	// Add Registered Routes
	mux := app.mount()

	log.Fatal(app.run(mux))
}
