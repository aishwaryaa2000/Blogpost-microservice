package main

import (
	"queryservice/router"
	"queryservice/apioperation"
	"fmt"
	"net/http"
)

func main(){
	handler:=router.CreateRoute()
	apioperation.ProcessPendingRequests()
	fmt.Println("Queryservice sever started at 4003")
	http.ListenAndServe(":4003",handler)
}

