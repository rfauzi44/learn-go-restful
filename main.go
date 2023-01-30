package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rfauzi44/learn-go-restful/router"
)

func main() {
	r, err := router.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	port := ":3000"
	log.Printf("server running on port %s", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}
