package router

import (
	"blogpost/apioperation"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	// "io"

)

func MuxRoute(){

// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)

	r := mux.NewRouter()
	r.HandleFunc("/eventbus/event/listener",apioperation.Post).Methods("POST","OPTIONS")
//	r.HandleFunc("/api/v1/blog/post",apioperation.GetPost).Methods("GET","OPTIONS") 

	
	// resp, _ := http.Get("http://localhost:4005/eventbus/event")
    // if resp.Body!=nil{
	// 	jsonData, _ := io.ReadAll(resp.Body)
	// 	fmt.Println(string(jsonData))
	// }

	//x=handle
	fmt.Println("Sever started")
	http.ListenAndServe(":4003",r)

	// return r;

	/* */
}