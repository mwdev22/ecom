package main

import (
	"log"

	"github.com/mwdev22/ecom/app/api"
)

func main() {

	server := api.NewServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
