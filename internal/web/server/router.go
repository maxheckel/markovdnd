package server

import (
	web "github.com/maxheckel/markovdnd/internal/web/handlers"
	"github.com/maxheckel/markovdnd/internal/web/handlers/get"
	"github.com/maxheckel/markovdnd/internal/web/handlers/run"
)

func (s server) AddRoutes(){

	s.Router.Handle("/run/chain/{name}", run.Chain{
		Store: s.Store,
	})
	s.Router.Handle("/get/crawled", get.Crawled{
		Store: s.Store,
	})
	spa := web.SpaHandler{StaticPath: "public", IndexPath: "index.html"}
	s.Router.PathPrefix("/").Handler(spa)
}