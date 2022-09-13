package router

import (
	"aes/apioperation"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateRoute() http.Handler{

	r := mux.NewRouter()
	r.HandleFunc("/aes/encrypt",apioperation.TripleAESEncrypt).Methods("POST")
	r.HandleFunc("/aes/decrypt",apioperation.TripleAESDecrypt).Methods("POST")

	return r
}