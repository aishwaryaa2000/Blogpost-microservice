package router

import (
	"blogpost/apioperation"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

)

func MuxRoute(){

// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/blog/post",apioperation.CreatePost).Methods("POST","OPTIONS")
	r.HandleFunc("/api/v1/blog/post",apioperation.GetPost).Methods("GET","OPTIONS") 

	
	fmt.Println("Sever started")
	http.ListenAndServe(":4001",r)

	// return r;

	/* */
}