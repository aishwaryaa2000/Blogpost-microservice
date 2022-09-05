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
	r.HandleFunc("/eventbus/event",apioperation.GetEvent).Methods("GET","OPTIONS")
	r.HandleFunc("/eventbus/event",apioperation.PostEvent).Methods("POST","OPTIONS") 

	
	fmt.Println("Sever started")
	http.ListenAndServe(":4005",r)

	// return r;

	/* */
}