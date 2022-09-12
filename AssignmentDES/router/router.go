package router

import (
	"des/apioperation"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateRoute() http.Handler{

	r := mux.NewRouter()
	r.HandleFunc("/des/encrypt",apioperation.TripleDESEncrypt).Methods("POST")
	r.HandleFunc("/des/decrypt",apioperation.TripleDESDecrypt).Methods("POST")

	return r
}