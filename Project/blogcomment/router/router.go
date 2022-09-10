package router

import (
	"blogcomment/apioperation"
	"net/http"
	"github.com/gorilla/mux"
	
)

func CreateRoute() http.Handler{

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/blog/post/{postId}/comment",apioperation.CreateComment).Methods("POST","OPTIONS")
	router.HandleFunc("/api/v1/blog/post/{postId}/comment",apioperation.GetComment).Methods("GET","OPTIONS")
	router.HandleFunc("/senteventafterprocess",apioperation.FinalEvent).Methods("POST","OPTIONS") 

	return router
	
}