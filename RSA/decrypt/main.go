package main

import(
	"net/http"
	"fmt"
	"decrypt/router"
) 

func main() {
	handler := router.CreateRoute()
	fmt.Println("Starting the server at 8081")
	http.ListenAndServe(":8081",handler)
}