package server

import "github.com/maxheckel/markovdnd/internal/web/handlers/run"

func (s server) AddRoutes(){
	s.Router.Handle("/run/chain/{name}", run.Chain{
		Store: s.Store,
	})
}