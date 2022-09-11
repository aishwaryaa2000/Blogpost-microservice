package apioperation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"blogcomment/model"
	"net/http"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)


var AllComments =make(map[uuid.UUID][]model.Comment) 
/*AllComments is a map of - postID : []{commentId,message}
example of a single entry is postId1 : [{cId1,message1}, {cId2,message2},{cId3,message3}...]*/

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	//this is used to avoid cors error
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func CreateComment(w http.ResponseWriter,r *http.Request){
	/*
	Here,the EventComment struct has the type as comment created
	and data interface will have commentId, message, id
	Example - {
		"type" : "comment created"
		"data" : {
			"id" : post ID here
			"commentId" : comment ID here
			"message" : comment msg here
		}
	}
	*/
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	jsonData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

   var receivedEvent model.EventComment
   json.Unmarshal(jsonData, &receivedEvent.Data)
   params := mux.Vars(r)
   paramPostIdUUID,_:=uuid.FromString(params["postId"])

   //Initialising the structure with appropriate values
   receivedEvent.Data["id"] =paramPostIdUUID
   receivedEvent.Type="Comment created"
   CommentId,_:=uuid.NewV4()
   receivedEvent.Data["commentId"] =CommentId

   CommentJsonEvent,_ := json.Marshal(&receivedEvent)
   fmt.Println("\nSending data from blogcomment to eventbus 4005 : ",string(CommentJsonEvent))

   //Sending post request to eventbus with requestbody as the event
   resp, err := http.Post("http://eventbus-serv:4005/eventbus/event", "application/json", bytes.NewBuffer(CommentJsonEvent))
   if err != nil {
		fmt.Println(err)
    }
   if resp.StatusCode==200{
		jsonresponse, _ := io.ReadAll(resp.Body)
		fmt.Println("Response in comment : ",string(jsonresponse))
		w.Write([]byte("Comment creation event cycle is completed"))
	}


}

func UpdateAllCommentsData(w http.ResponseWriter,r *http.Request){
	/*
	FINAL STEP OF THE EVENT CYCLE FOR COMMENT CREATION
	After recieving the processed event from the event bus in the request body,
	we update the data storage of the blogcomments
	here,we append the newly created comment to the postID entry in the map AllComments
	*/
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}

	var processedEvent model.EventComment
	processedEventJson, _ := io.ReadAll(r.Body)
	json.Unmarshal(processedEventJson,&processedEvent)
	fmt.Println("This is the event we recieved in blogcomment by event bus in request body: ",string(processedEventJson))

	//extracting the data from the event E recieved in request body
	postIdUUID := uuid.Must(uuid.FromString(processedEvent.Data["id"].(string)))	
	newComment := model.Comment{
		CommentId: uuid.Must(uuid.FromString(processedEvent.Data["commentId"].(string))),
		Message: processedEvent.Data["message"].(string),
	}

	//Updating the data storage of blogcomment i.e Allcomments
	allCommentsOfId := AllComments[postIdUUID]
	allCommentsOfId = append(allCommentsOfId,newComment)
	AllComments[postIdUUID]= allCommentsOfId

	fmt.Println("All comments of ID ",postIdUUID," :",allCommentsOfId)
	w.Write([]byte("Successfully got event from EB to blogcomment after all the processing"))

}


func GetCommentsWithPostId(w http.ResponseWriter,r *http.Request){
	/*This handle function sends all the comments of a particular post ID specified in the URL
	as response body whenever the appropriate endpoint is hit via frontend*/

	setupCorsResponse(&w,r)
	params := mux.Vars(r)
	paramPostId := params["postId"] 
	paramPostIdUUID,_:=uuid.FromString(paramPostId)

	fmt.Println("All comments : ",AllComments[paramPostIdUUID])
	json.NewEncoder(w).Encode(AllComments[paramPostIdUUID])

}