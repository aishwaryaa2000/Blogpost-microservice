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


var posts =make(map[uuid.UUID]*PostResponseData) 


type PostResponseData struct{
	Id uuid.UUID `json:"id"`
	Comments map[uuid.UUID]string `json:"comments"`
	Title string `json:"title"`
	// Type string `json:"type"`
}

// type Comment struct{
// 	postId uuid.UUID `json:"postId"`
// 	commentId uuid.UUID	`json:"commentId"`
// 	msg string

// }

type PostRequestData struct{
	postId uuid.UUID `json:"id"`
	commentId uuid.UUID	`json:"id"`
	msg string `json:"message"`
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
	requestDataJson, err := ioutil.ReadAll(r.Body)
	//requestData in json
	if err != nil {
		log.Fatal(err)
	}
	var requestPostData PostRequestData
   	json.Unmarshal(requestDataJson,&requestPostData)
	fmt.Println("Event recieved : ",requestPostData.Type)
	//data has type id title of the post
	singlePost:=handleEvent(&requestPostData)
	json.NewEncoder(w).Encode(singlePost)
}

func handleEvent(request *PostRequestData) PostResponseData{
	if (request.Type == "Post Created"){
		fmt.Println(request)
		var newPost PostResponseData
		var temp = make(map[uuid.UUID]string)
		newPost.Comments=temp; //null map
		newPost.Id=request.postId
		newPost.Title=request.Title
		posts[request.postId] = &newPost
        return *posts[request.postId];
    }
    //for comments we need the comment ID ,message and postID
    // if (request.Type == "Comment Created"){
		fmt.Println(request)
		posts[request.postId].Comments[request.commentId]=request.msg
        return *posts[request.postId];
    // }
}

func GetPost(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	fmt.Println("Helloo")
	json.NewEncoder(w).Encode(posts)

}