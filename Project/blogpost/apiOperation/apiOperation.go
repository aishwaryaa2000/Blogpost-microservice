package apioperation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gofrs/uuid"
	"blogpost/model"
)

var posts =make([]*model.Post,0) 

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
	jsonData, _ := io.ReadAll(r.Body)
	
	var receivedEvent model.EventPost
	json.Unmarshal(jsonData, &receivedEvent.Data)
	receivedEvent.Type="Post created"
	postId,_:=uuid.NewV4()
	receivedEvent.Data["id"] =postId
	PostJsonEvent,_ := json.Marshal(&receivedEvent)

	fmt.Println("\nSending data from blogpost to eventbus 4005 : ",string(PostJsonEvent))
	resp, err := http.Post("http://localhost:4005/eventbus/event", "application/json",bytes.NewBuffer(PostJsonEvent))
	if err != nil {
		fmt.Println("error in blogpost : ",err)
	}

	if resp.StatusCode==200{
		jsonresponse, _ := io.ReadAll(resp.Body)
		fmt.Println("Response in blogpost : ",string(jsonresponse))
		w.Write([]byte("Post creation event cycle is completed"))
	}
		
}

func FinalEvent(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
   	var processedEvent model.EventPost
	processedEventJson, _ := io.ReadAll(r.Body)
	json.Unmarshal(processedEventJson,&processedEvent)
	
	fmt.Println("This is the event we recieved in blogcomment by event bus in request body: ",string(processedEventJson))

	postIdUUID := uuid.Must(uuid.FromString(processedEvent.Data["id"].(string)))

	newPost := model.Post{
		Id: postIdUUID,
		Title: processedEvent.Data["title"].(string),
	}

	fmt.Println("New post is : ",newPost)
	posts=append(posts, &newPost)
	w.Write([]byte("Successfully got event from EB to blogpost after all the processing"))

}

func GetPost(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	json.NewEncoder(w).Encode(posts)

}