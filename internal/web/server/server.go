package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/maxheckel/markovdnd/internal/services/store/chain"
	"net/http"
)

type Server interface{
	Start() error

}

type server struct {
	Router *mux.Router
	Store chain.Store
}


func (s server) Start() error {
	s.SetStore()
	s.AddRoutes()
	port := ":80"
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})


	fmt.Printf("Starting server on %s", port)
	http.Handle("/", handlers.CORS(headersOk, originsOk, methodsOk)(s.Router))

	return http.ListenAndServe(port, nil)
}



func New() Server {
	s := server{}
	s.Router = mux.NewRouter()
	return s
}