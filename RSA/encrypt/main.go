package main

import(
	"net/http"
	"fmt"
	"encrypt/router"
) 

func main() {
	handler := router.CreateRoute()
	fmt.Println("Starting the server at 8080")
	http.ListenAndServe(":8080",handler)
}