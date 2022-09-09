package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"queryservice/apioperation"
	// "bytes"
	"github.com/gorilla/mux"
)

func MuxRoute(){

// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	var queue []apioperation.PostRequestData
	r := mux.NewRouter()
	r.HandleFunc("/eventbus/event/listener",apioperation.Post).Methods("POST","OPTIONS")
//	r.HandleFunc("/api/v1/blog/post",apioperation.GetPost).Methods("GET","OPTIONS") 

	
	resp, _ := http.Get("http://localhost:4005/eventbus/event/queue")
    if resp!=nil{
		jsonData, _ := io.ReadAll(resp.Body)
		json.Unmarshal(jsonData,&queue)
		fmt.Println(queue)
	}


	for _,singleEvent := range queue{
		//jsonSingleEvent,_ := json.Marshal(singleEvent)
		apioperation.ProcessEvent(singleEvent)
		// r,_ := http.NewRequest("POST","/eventbus/event/listener",bytes.NewBuffer(jsonSingleEvent))
		// apioperation.Post(http.ResponseWriter, r)


	}


	


	//x=handle
	fmt.Println("Sever started")
	http.ListenAndServe(":4003",r)

	// return r;

	/* */
}