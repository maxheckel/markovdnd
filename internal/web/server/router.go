package server

import "github.com/maxheckel/auto-dnd/internal/web/handlers/run"

func (s server) AddRoutes(){
	s.Router.Handle("/run/chain", run.Chain{})
}