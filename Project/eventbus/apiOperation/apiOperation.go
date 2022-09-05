package apioperation

import (
	"bytes"
	// "encoding/json"
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "os"
	"github.com/gofrs/uuid"
)

// var posts =make(map[uuid.UUID]*Post)

// type Post struct{
// 	PostId uuid.UUID `json:"id"`
// 	Title string `json:"title"`
// }


type Post struct{
	Type string `json:"type"`
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`

}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func PostEvent(w http.ResponseWriter,r *http.Request){
	//iske request body me we have title type and id
	//create post event	
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}

	jsonData, _ := io.ReadAll(r.Body)
	fmt.Println("while sending r body to query service : ",string(jsonData))

	//requestData in json

	// if err != nil {
	// 	log.Fatal(err)
	// }
   
	//send the json object to the eventbus
	fmt.Println("\nSending data from event bus to queryservice 4003 ",r.Body)
	resp, err := http.Post("http://localhost:4003/eventbus/event/listener", "application/json",bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode==200{
		jsonData, _ := io.ReadAll(resp.Body)

		fmt.Println(string(jsonData))
		//resp.body should have id and title
		w.Write(jsonData)
	}

		
}

func GetEvent(w http.ResponseWriter,r *http.Request){
	
}