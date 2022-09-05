package apioperation

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	//  "bytes"
	// "os"
	"github.com/gofrs/uuid"


)


var posts =make(map[uuid.UUID]*PostResponseData) 


type PostResponseData struct{
	Id uuid.UUID `json:"id"`
	Comments map[uuid.UUID]string `json:"comments"`
	Title string `json:"title"`
	// Type string `json:"type"`
}

// type Comment struct{
// 	Id uuid.UUID `json:"Id"`
// 	commentId uuid.UUID	`json:"commentId"`
// 	msg string
// }

type PostRequestData struct{
	Id uuid.UUID `json:"id"`
	CommentId uuid.UUID	`json:"commentId"`
	Msg string `json:"message"`
	Type string `json:"type"`
	Title string `json:"title"`
	// Comments map[uuid.UUID]string `json:"comments"`

}


func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func Post(w http.ResponseWriter,r *http.Request){
	
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	requestDataJson, err := io.ReadAll(r.Body)
	//requestData in json
	fmt.Println("In string : ",string(requestDataJson))
	if err != nil {
		log.Fatal(err)
	}
	var requestPostData PostRequestData
   	json.Unmarshal(requestDataJson,&requestPostData)
	fmt.Println("Event recieved : ",requestPostData.Type)
	//data has type id title of the post
	fmt.Println("INSIDE QUERY : ",requestPostData)
	singlePost:=handleEvent(&requestPostData)
	fmt.Println(singlePost)
	json.NewEncoder(w).Encode(singlePost)
}

func handleEvent(request *PostRequestData) PostResponseData{
	if (request.Type == "Post created"){
		var newPost PostResponseData
		var temp = make(map[uuid.UUID]string)
		newPost.Comments=temp; //null map
		newPost.Id=request.Id
		newPost.Title=request.Title
		fmt.Println("\nThis is new post inside query : ",newPost)

		posts[request.Id] = &newPost
        return *posts[request.Id];
    }
    //for comments we need the comment ID ,message and Id
    // if (request.Type == "Comment Created"){
		fmt.Println(request)
		posts[request.Id].Comments[request.CommentId]=request.Msg
        return *posts[request.Id];
    // }
}

func GetPost(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	fmt.Println("Helloo")
	json.NewEncoder(w).Encode(posts)

}