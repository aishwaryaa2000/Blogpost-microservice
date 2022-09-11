package router

import (
	"eventbus/apioperation"
	"net/http"
	"github.com/gorilla/mux"

)

func CreateRoute() http.Handler{

	router := mux.NewRouter()
	router.HandleFunc("/eventbus/event/queue",apioperation.SendEventQueue).Methods("GET","OPTIONS")
	router.HandleFunc("/eventbus/event",apioperation.PostEvent).Methods("POST","OPTIONS") 
	router.HandleFunc("/eventbus/processedevent",apioperation.RedirectProcessedEvent).Methods("POST","OPTIONS") 

	return router

}