//This will get a json object
package main

import (
	"net/http"
	"blogcomment/router"
	"fmt"
)

func main(){
	handler:=router.CreateRoute()
	fmt.Println("Blog comment Sever started at 4002")
	http.ListenAndServe(":4002",handler)
}
