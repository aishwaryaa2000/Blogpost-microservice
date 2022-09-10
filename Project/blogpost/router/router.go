package router

import (
	"blogpost/apioperation"
	"net/http"
	"github.com/gorilla/mux"

)

func CreateRoute() http.Handler{

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/blog/post",apioperation.CreatePost).Methods("POST","OPTIONS")
	router.HandleFunc("/api/v1/blog/post",apioperation.GetPost).Methods("GET","OPTIONS") 
	router.HandleFunc("/senteventafterprocess",apioperation.FinalEvent).Methods("POST","OPTIONS") 

	return router;

}

