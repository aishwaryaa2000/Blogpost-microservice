package apioperation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "os"
	"github.com/gofrs/uuid"
)


var posts =make(map[uuid.UUID]*Post) 


// type Post struct{
// 	PostId uuid.UUID `json:"id"`
// 	Title string `json:"title"`
// }


type Post struct{
	Type string `json:"type"`
	PostId uuid.UUID `json:"id"`
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

	jsonData, err := io.ReadAll(r.Body)
	//requestData in json

	if err != nil {
		log.Fatal(err)
	}
   
	//send the json object to the eventbus
	resp, err := http.Post("https://localhost:4003/eventbus/event/listener", "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode==200{
		fmt.Println(resp.Body)
		//resp.body should have id and title
		json.NewEncoder(w).Encode(resp.Body)
	}

		
}

func GetEvent(w http.ResponseWriter,r *http.Request){
	
}