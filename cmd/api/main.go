package main

import (
	"log"

	"github.com/mahmudul-hasan01/social-media-backend-golang/internal/env"
	"github.com/mahmudul-hasan01/social-media-backend-golang/internal/store"
)

func main() {

	cfg :=  config{
			addr: env.GetString("ADDR", ":8080"),
		}

		store := store.NewPostgresStorage(nil)
		
	app := &application{
		config: cfg,
		store: store,
	}


	mux := app.mount()

	log.Fatal(app.run(mux))
		
}