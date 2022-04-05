package server

import (
	"github.com/maxheckel/markovdnd/internal/web/handlers/get"
	"github.com/maxheckel/markovdnd/internal/web/handlers/run"
	"net/http"
)

func (s server) AddRoutes(){
	s.Router.Handle("/run/chain/{name}", run.Chain{
		Store: s.Store,
	})
	s.Router.Handle("/get/crawled", get.Crawled{
		Store: s.Store,
	})
	fs := http.FileServer(http.Dir("./public/"))
	s.Router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
}