package router

import (
	"blogcomment/apioperation"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

)

func MuxRoute(){

// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/blog/post/{postId}/comment",apioperation.CreateComment).Methods("POST","OPTIONS")
	r.HandleFunc("/api/v1/blog/post/{postId}/comment",apioperation.GetComment).Methods("GET","OPTIONS")

	// r.HandleFunc("/api/v1/blog/post",apioperation.GetPost).Methods("GET","OPTIONS") 

	// await axios.post(`http://localhost:4002/api/v1/blog/post/${postId}/comment`,{message}).catch(e=>console.log(e.message))

	fmt.Println("Blog comment Sever started at 4002")
	http.ListenAndServe(":4002",r)

	// return r;

	/* */
}