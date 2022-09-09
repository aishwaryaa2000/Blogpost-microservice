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
   data.Id,_=uuid.NewV4()
   data.Type="Post created"
   fmt.Println("Data in blogpost is : ",data)
   PostDataJson,_:=json.Marshal(data)
   
	//send the json object to the eventbus
	fmt.Println("\nSending data from blogpost to eventbus 4005")
	resp, err := http.Post("http://localhost:4005/eventbus/event", "application/json",bytes.NewBuffer(PostDataJson))
	if err != nil {
		fmt.Println("error in blogpost : ",err)
	}

	if resp.StatusCode==200{
		jsonresponse, _ := io.ReadAll(resp.Body)
		fmt.Println("Response in blogpost : ",string(jsonresponse))
	}
		
}

func FinalEvent(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	//requestData in json
   	var getResponseFromEventBus Post
	responseFromEventBus, _ := io.ReadAll(r.Body)
	json.Unmarshal(responseFromEventBus,&getResponseFromEventBus)

	fmt.Println("This is the r body we recieved in blogpost by event bus : ")
	fmt.Println(getResponseFromEventBus)
	posts[getResponseFromEventBus.Id]=&getResponseFromEventBus
	//send first response from bp to eb send OK
	w.Write([]byte("Successfully got event from EB to blogpost after all the processing"))

}

func GetPost(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	fmt.Println("Helloo")
	json.NewEncoder(w).Encode(posts)

}