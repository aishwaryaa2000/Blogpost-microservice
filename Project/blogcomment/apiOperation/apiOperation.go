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
	jsonData, err := io.ReadAll(r.Body)
	//requestData in json

	if err != nil {
		log.Fatal(err)
	}

   var receivedEvent model.EventComment
   json.Unmarshal(jsonData, &receivedEvent.Data)

   params := mux.Vars(r)
   paramPostIdUUID,_:=uuid.FromString(params["postId"])
   receivedEvent.Data["id"] =paramPostIdUUID

   receivedEvent.Type="Comment created"

   CommentId,_:=uuid.NewV4()
   receivedEvent.Data["commentId"] =CommentId

   fmt.Println(receivedEvent) 

   CommentJsonEvent,_ := json.Marshal(&receivedEvent)

   fmt.Println("\nSending data from blogcomment to eventbus 4005")
   resp, err := http.Post("http://localhost:4005/eventbus/event", "application/json", bytes.NewBuffer(CommentJsonEvent))
   if err != nil {
		fmt.Println(err)
    }
   if resp.StatusCode==200{
		jsonresponse, _ := io.ReadAll(resp.Body)
		fmt.Println("Response in comment : ",string(jsonresponse))
		w.Write([]byte("Comment creation event cycle is completed"))
	}


}

func FinalEvent(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	//requestData in json
	var processedEvent model.EventComment
	processedEventJson, _ := io.ReadAll(r.Body)
	//This is the response from event bus
	json.Unmarshal(processedEventJson,&processedEvent)

	fmt.Println("This is the event we recieved in blogcomment by event bus in request body: ",string(processedEventJson))

	postIdUUID := uuid.Must(uuid.FromString(processedEvent.Data["id"].(string)))
	
	newComment := model.Comment{
		CommentId: uuid.Must(uuid.FromString(processedEvent.Data["commentId"].(string))),
		Message: processedEvent.Data["message"].(string),
	}
	allCommentsOfId := AllComments[postIdUUID]
	allCommentsOfId = append(allCommentsOfId,newComment)
	AllComments[postIdUUID]= allCommentsOfId

	fmt.Println("all comments of that ID : " ,allCommentsOfId)

	w.Write([]byte("Successfully got event from EB to blogcomment after all the processing"))

}


func GetComment(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	params := mux.Vars(r)
	paramPostId := params["postId"] //we got the post id
	paramPostIdUUID,_:=uuid.FromString(paramPostId)

	fmt.Println("all : ",AllComments[paramPostIdUUID])
	json.NewEncoder(w).Encode(AllComments[paramPostIdUUID])

}