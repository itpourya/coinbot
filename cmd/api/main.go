package main

import (
	"CoinBot/internal/server"
	"fmt"
	"log"
)

func main() {

	srv := server.NewServer()

	log.Println("Server started")
	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
