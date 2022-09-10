package main

import (
	"fmt"
	"net/http"
	"eventbus/router"

)

func main(){
	handler:=router.CreateRoute()
	fmt.Println("Eventbus sever started at 4005")
	http.ListenAndServe(":4005",handler)
	
}