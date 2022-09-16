package router

import (
	"encrypt/apioperation"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateRoute() http.Handler{
	r := mux.NewRouter()
	r.HandleFunc("/rsa/encrypt",apioperation.SendMsg).Methods("POST")
	
	return r
}