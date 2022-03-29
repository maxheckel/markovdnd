package main

import (
	"github.com/maxheckel/auto-dnd/internal/web/server"
	"log"
)

func main() {
	srv := server.New()

	log.Fatalln(srv.Start())
}