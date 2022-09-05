package apioperation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)


var posts =make(map[uuid.UUID]*Post) 


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

func CreatePost(w http.ResponseWriter,r *http.Request){
	
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	jsonData, err := io.ReadAll(r.Body)
	//requestData in json

	if err != nil {
		log.Fatal(err)
	}
   var data Post
   json.Unmarshal(jsonData, &data)
   fmt.Println("Data in blogpost is : ",data)
   data.Id,_=uuid.NewV4()
   data.Type="Post created"
   PostDataJson,_:=json.Marshal(data)
   
	//send the json object to the eventbus
	fmt.Println("\nSending data from blogpost to eventbus 4005")
	resp, err := http.Post("http://localhost:4005/eventbus/event", "application/json",
		bytes.NewBuffer(PostDataJson))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode==200{
		// var res map[string]interface{}
		//response body of event bus will give me id and title
		var getResponseFromEventBus Post
		responseFromEventBus, _ := io.ReadAll(resp.Body)
		json.Unmarshal(responseFromEventBus,&getResponseFromEventBus)

		fmt.Println("This is the response we recieved in blogpost by event bus : ")
		fmt.Println(getResponseFromEventBus)
		posts[getResponseFromEventBus.Id]=&getResponseFromEventBus
	}
		
}

func GetPost(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	fmt.Println("Helloo")
	json.NewEncoder(w).Encode(posts)

}