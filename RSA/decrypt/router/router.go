package router

import (
	"decrypt/apioperation"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateRoute() http.Handler{
	r := mux.NewRouter()
	r.HandleFunc("/rsa/publickey",apioperation.SendPublicKey).Methods("GET")
	r.HandleFunc("/rsa/decrypt",apioperation.GeneratePlaintext).Methods("POST")

	return r
}