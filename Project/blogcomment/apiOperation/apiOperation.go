package apioperation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"

	// "os"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)



var AllComments =make(map[uuid.UUID][]Comment) 

type Comment struct{
	CommentId uuid.UUID `json:"commentId"`
	Message string `json:"message"`
}


type ResponseFromEventBus struct{
	Comments map[uuid.UUID]string `json:"comments"`
	Id uuid.UUID `json:"id"`
}


type CommentRequestObject struct{
	Type string `json:"type"`
	Id uuid.UUID `json:"id"`
	CommentId uuid.UUID `json:"commentId"`
	Message string `json:"message"`
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func CreateComment(w http.ResponseWriter,r *http.Request){
	
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	requestData, err := io.ReadAll(r.Body)
	//requestData in json

	if err != nil {
		log.Fatal(err)
	}
	params := mux.Vars(r)
	paramPostId := params["postId"] //we got the post id


   var requestObject CommentRequestObject
   requestObject.CommentId,_=uuid.NewV4()
   fmt.Println("comment id is : ",requestObject.CommentId)
   paramPostIdUUID,_:=uuid.FromString(paramPostId)
   requestObject.Id= paramPostIdUUID
   requestObject.Type = "Comment created"
   json.Unmarshal(requestData, &requestObject)
   fmt.Println(requestObject) 
   PostDataJson,_:=json.Marshal(requestObject)


   fmt.Println("\nSending data from blogcomment to eventbus 4005")
   resp, err := http.Post("http://localhost:4005/eventbus/event", "application/json",
	   bytes.NewBuffer(PostDataJson))
   if err != nil {
	   log.Fatal(err)
   }
	 
   if resp.StatusCode==200{
	// var res map[string]interface{}
	//response body of event bus will give me id and title
	var getResponseFromEventBus ResponseFromEventBus
	responseFromEventBus, _ := io.ReadAll(resp.Body)
	fmt.Println("ISNIDE BLOG COMMENT response from eventbus : ",string(responseFromEventBus))
	json.Unmarshal(responseFromEventBus,&getResponseFromEventBus)

	fmt.Println("This is the response we recieved in blogcomment by event bus : ")
	fmt.Println(getResponseFromEventBus.Comments)
	// temp:=AllComments[getResponseFromEventBus.Id]
	var temp []Comment
	for key, element := range getResponseFromEventBus.Comments {
        fmt.Println("Key:", key, "=>", "Element:", element)
		var singleComment Comment
		singleComment.CommentId=key
		singleComment.Message=element
		temp=append(temp,singleComment)
    }
	AllComments[getResponseFromEventBus.Id]= temp
	fmt.Println(AllComments)
	fmt.Println(reflect.TypeOf(getResponseFromEventBus.Comments))
	// AllComments=getResponseFromEventBus.Comments	
 	//AllComments[getResponseFromEventBus.Id]=&getResponseFromEventBus
}
	//    var commentRequestObject CommentRequestObject;
	//    commentRequestObject.PostId=requestObject.PostId;
	//    commentRequestObject.Title=requestObject.Title;
	//    commentRequestObject.Type = "Comment created"
	//    postRequestDataJson, _ := json.Marshal(postRequestObject)
	//    //send the json object to the eventbus
	// 	resp, err := http.Post("https://localhost:4005/eventbus/event", "application/json",
	// 		bytes.NewBuffer(postRequestDataJson))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}




	// 	var res map[string]interface{}
	// 	json.NewDecoder(resp.Body).Decode(&res)
	// 	fmt.Println(res["json"])
}

func GetComment(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	params := mux.Vars(r)
	paramPostId := params["postId"] //we got the post id
	paramPostIdUUID,_:=uuid.FromString(paramPostId)

	fmt.Println("all : ",AllComments[paramPostIdUUID])
	json.NewEncoder(w).Encode(AllComments[paramPostIdUUID])

}