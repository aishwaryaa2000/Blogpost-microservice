package main

import (
	"hashintegrity/router"
	"fmt"
	"net/http"
)


func main(){

	handler := router.CreateRoute()
	fmt.Println("Sever started at 8081")
	http.ListenAndServe(":8081",handler)	
}
