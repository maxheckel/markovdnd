package main

import (
	"github.com/maxheckel/markovdnd/internal/web/server"
	"log"
)

func main() {
	srv := server.New()

	log.Fatalln(srv.Start())
}