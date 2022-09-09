//This will get a json object
package main

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "bytes"
	// "os"
	"blogpost/router"


)





func main(){
	router.MuxRoute()
	//router.HandleFunc("/<your-url>", <function-name>).Methods("<method>")
	
	//post and get
	// POST is used to send data to a server to create/update a resource. 
	// The data sent to the server with POST is stored in the request body of the HTTP request
	// fmt.Println("Sever started")
	// http.ListenAndServe(":4001",r)

	
}




















// type Posts struct{
// 	U []Post `json:"Posts"`
// }






// type PostRequest struct{
// 	Type string `json:"type"`
// 	PostId string `json:"id"`
// 	Title string `json:"title"`
// }


	   	// POST is used to send data to a server to create/update a resource. 
	// The data sent to the server with POST is stored in the request body of the HTTP request
	// r is request from client to server


	//response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	//r is request from client to server
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// 	}
	// responseData, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//    var responseObject Post
	//    json.Unmarshal(responseData, &responseObject)

	//    fmt.Println(responseObject.PostId)
	//    fmt.Println(responseObject.Title)




	
