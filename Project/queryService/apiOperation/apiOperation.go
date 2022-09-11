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
//posts consists of all the posts along with id,title and []comments{cId,message}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	//to avoid CORS error
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}


func Post(w http.ResponseWriter,r *http.Request){
	//this handle func sends the event to handleEvent() wherein data will be added to "posts"
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
	/*we check the type of the event and
	  store the newly created post or newly created comment 
	  into the data storage of queryservice i.e "posts"
	*/
	dataMap := receivedEvent.Data
	postIdUUID := uuid.Must(uuid.FromString(dataMap["id"].(string)))

	if (receivedEvent.Type == "Post created"){
		//for a new post,the comments should be empty
		var comments = make(map[uuid.UUID]string)
		newPost := model.SinglePost{
			Id: postIdUUID,
			Comments: comments,
			Title: dataMap["title"].(string),

		}
		fmt.Println("\nThis is new post inside query service : ",newPost)

		//updating data storage of querservice
		posts[postIdUUID] = &newPost

    }else{
		//if type is comment created then data interface has the postID,CommentID and the message of the comment
		_,ok:=posts[postIdUUID]
		fmt.Println("Post with this key exists? : ",ok)
		/*since data "posts" in QS is not persistent whenever QS goes down
		so first, we need to check if post with that ID exists or not
		ideally,in such cases we have to use local storage property of the browser or durability property of a database.
		But for now,we will create a new post with that ID and title as new post.*/
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

		//updating data storage of querservice
		commentIdUUID := uuid.Must(uuid.FromString(dataMap["commentId"].(string)))
		posts[postIdUUID].Comments[commentIdUUID]=dataMap["message"].(string)

	}

	fmt.Println("All posts : ",posts)

	//send the event to eventBus once the data storage of QS is updated
	sendEventToEB(receivedEvent)

}

func sendEventToEB(receivedEvent model.Event){
	//this func() sends the event to eventBus
	fmt.Println("\n\nAfter successfully updating QS db\nSending event from query service to eventbus after processing : ",receivedEvent)
	eventJson,_ := json.Marshal(receivedEvent)
	resp, err := http.Post("http://eventbus-serv:4005/eventbus/processedevent", "application/json",bytes.NewBuffer(eventJson))

	if resp.StatusCode==200{
		responseJson, _ := io.ReadAll(resp.Body)
		fmt.Println("This is the response of post method sent from EB to QS : ",string(responseJson))
	}
	if err!=nil{
		fmt.Println("\nError in QS is : ",err)
	}

}


func ProcessPendingRequests(){
	//this func() is used to process the events which were initiated when QS was down
	var queue []model.Event
	resp, _ := http.Get("http://eventbus-serv:4005/eventbus/event/queue")
    if resp!=nil{
		jsonData, _ := io.ReadAll(resp.Body)
		json.Unmarshal(jsonData,&queue)
		fmt.Println(queue)
	}

	for _,singleEvent := range queue{
		HandleEvent(singleEvent)
	}
}