package apioperation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"
	"github.com/gofrs/uuid"
	"queryservice/model"
)

var posts =make(map[uuid.UUID]*model.SinglePost) 

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
	eventJson, _ := io.ReadAll(r.Body)
	fmt.Println("Event got from EB to QS : ",string(eventJson))
	var receivedEvent model.Event
   	json.Unmarshal(eventJson,&receivedEvent)
	HandleEvent(receivedEvent)

	w.Write([]byte("OK from QS to EB.Event received and processed by QS"))

}

func HandleEvent(receivedEvent model.Event) {
	dataMap := receivedEvent.Data
	postIdUUID := uuid.Must(uuid.FromString(dataMap["id"].(string)))

	if (receivedEvent.Type == "Post created"){
		var comments = make(map[uuid.UUID]string)

		newPost := model.SinglePost{
			Id: postIdUUID,
			Comments: comments,
			Title: dataMap["title"].(string),

		}
		fmt.Println("\nThis is new post inside query : ",newPost)

		posts[postIdUUID] = &newPost

    }else{
		//if type is comment created then data interface has the postID,CommentID and the message of the comment
		_,ok:=posts[postIdUUID]
		fmt.Println("Post with this key exists? : ",ok)
		/*since data "posts" in QS is not persistent so we need to check if post with that ID exists or not
		in such cases we have to use local storage property of the browser or durability property of a database.
		But for now,we will create a new post with that ID.*/
		if (!ok){
				//if we lost the data of QS then create a new post and then append the comments
				var comments = make(map[uuid.UUID]string)
				newPost := model.SinglePost{
					Id: postIdUUID,
					Comments: comments,
					Title: "new post",
				}

				posts[postIdUUID] = &newPost
		}

		commentIdUUID := uuid.Must(uuid.FromString(dataMap["commentId"].(string)))
		posts[postIdUUID].Comments[commentIdUUID]=dataMap["message"].(string)

	}

	fmt.Println("All posts : ",posts)

	sendEventToEB(receivedEvent)

}

func sendEventToEB(receivedEvent model.Event){
	fmt.Println("\n\nAfter successfully updating QS db\nSending event from query service to eventbus after processing : ",receivedEvent)
	eventJson,_ := json.Marshal(receivedEvent)
	resp, err := http.Post("http://localhost:4005/eventbus/getevent", "application/json",bytes.NewBuffer(eventJson))

	if resp.StatusCode==200{
		responseJson, _ := io.ReadAll(resp.Body)
		fmt.Println("This is the response of post method sent from EB to QS : ",string(responseJson))
	}
	if err!=nil{
		fmt.Println("\nError in QS is : ",err)
	}

}


func ProcessPendingRequests(){
	var queue []model.Event

	resp, _ := http.Get("http://localhost:4005/eventbus/event/queue")
    if resp!=nil{
		jsonData, _ := io.ReadAll(resp.Body)
		json.Unmarshal(jsonData,&queue)
		fmt.Println(queue)
	}

	for _,singleEvent := range queue{
		HandleEvent(singleEvent)
	}
}