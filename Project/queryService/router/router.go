package router

import (
	"net/http"
	"queryservice/apioperation"
	"github.com/gorilla/mux"
)

func CreateRoute() http.Handler{

	router := mux.NewRouter()
	router.HandleFunc("/eventbus/event/listener",apioperation.Post).Methods("POST","OPTIONS")	
	return router
}
