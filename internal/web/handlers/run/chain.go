package run

import "net/http"

type Chain struct {}

func (rc Chain) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is my home page"))
}