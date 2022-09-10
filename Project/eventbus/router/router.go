package router

import (
	"eventbus/apioperation"
	"net/http"
	"github.com/gorilla/mux"

)

func CreateRoute() http.Handler{

	router := mux.NewRouter()
	router.HandleFunc("/eventbus/event/queue",apioperation.GetEventQueue).Methods("GET","OPTIONS")
	router.HandleFunc("/eventbus/event",apioperation.PostEvent).Methods("POST","OPTIONS") 
	router.HandleFunc("/eventbus/getevent",apioperation.AfterProcessEventFromQS).Methods("POST","OPTIONS") 

	return router

}