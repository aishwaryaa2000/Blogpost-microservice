package apioperation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	 "bytes"
	// "os"
	"github.com/gofrs/uuid"


)


var posts =make(map[uuid.UUID]*Post) 


type Post struct{
	PostId uuid.UUID `json:"id"`
	Title string `json:"title"`
}


type PostRequest struct{
	Type string `json:"type"`
	PostId uuid.UUID `json:"id"`
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
	requestData, err := ioutil.ReadAll(r.Body)
	//requestData in json

	if err != nil {
		log.Fatal(err)
	}
   var requestObject Post
   requestObject.PostId,_=uuid.NewV4()
   json.Unmarshal(requestData, &requestObject)

   if requestObject.Title!=""{
	fmt.Println("added",requestObject);
   	posts[requestObject.PostId] = &requestObject
   }
	  fmt.Println(requestObject)


	  
	   var postRequestObject PostRequest;
	   postRequestObject.PostId=requestObject.PostId;
	   postRequestObject.Title=requestObject.Title;
	   postRequestObject.Type = "Post created"
	   postRequestDataJson, _ := json.Marshal(postRequestObject)
	   //send the json object to the eventbus
		resp, err := http.Post("https://localhost:4005/eventbus/event", "application/json",
			bytes.NewBuffer(postRequestDataJson))
		if err != nil {
			log.Fatal(err)
		}




		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Println(res["json"])
}

func GetPost(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	fmt.Println("Helloo")
	json.NewEncoder(w).Encode(posts)

}