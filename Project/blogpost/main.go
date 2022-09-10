package main

import (
	"fmt"
	"net/http"
	"blogpost/router"

)

func main(){
	handler := router.CreateRoute()
	fmt.Println("Blog post sever started at 4001")
	http.ListenAndServe(":4001",handler)
	
}




