package router

import (
	"hashintegrity/apioperation"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateRoute() http.Handler{

	r := mux.NewRouter()
	r.HandleFunc("/md5/create",apioperation.SendMsgHash).Methods("POST")
	r.HandleFunc("/md5/check",apioperation.CheckMsgHash).Methods("POST")

	return r
}