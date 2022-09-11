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

var allPosts =make([]*model.Post,0) 
/*allPosts is a slice of Post struct
example [{id1,title1},{id2,title2},{id3,title3}...]*/


func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	//this is used to avoid cors error
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

	//Initialising the structure with appropriate values
	receivedEvent.Type="Post created"
	postId,_:=uuid.NewV4()
	receivedEvent.Data["id"] =postId

	PostJsonEvent,_ := json.Marshal(&receivedEvent)
	fmt.Println("\nSending data from blogpost to eventbus 4005 : ",string(PostJsonEvent))

	//Sending post request to eventbus with requestbody as the event
	resp, err := http.Post("http://eventbus-serv:4005/eventbus/event", "application/json",bytes.NewBuffer(PostJsonEvent))
	if err != nil {
		fmt.Println("error in blogpost : ",err)
	}

	if resp.StatusCode==200{
		jsonresponse, _ := io.ReadAll(resp.Body)
		fmt.Println("Response in blogpost : ",string(jsonresponse))
		w.Write([]byte("Post creation event cycle is completed"))
	}
		
}

func UpdateAllPostsData(w http.ResponseWriter,r *http.Request){
	/*
	FINAL STEP OF THE EVENT CYCLE FOR POST CREATION
	After recieving the processed event from the event bus in the request body,
	we update the data storage of the blogpost
	here,we append the newly created post{id,title} to the allPosts[]
	*/
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
   	var processedEvent model.EventPost
	processedEventJson, _ := io.ReadAll(r.Body)
	json.Unmarshal(processedEventJson,&processedEvent)
	fmt.Println("This is the event we recieved in blopost by event bus in request body: ",string(processedEventJson))

	//extracting the data from the event E recieved in request body
	postIdUUID := uuid.Must(uuid.FromString(processedEvent.Data["id"].(string)))
	newPost := model.Post{
		Id: postIdUUID,
		Title: processedEvent.Data["title"].(string),
	}
	fmt.Println("New post is : ",newPost)

	//Updating the data storage of blogpost i.e allPosts
	allPosts=append(allPosts, &newPost)

	w.Write([]byte("Successfully got event from EB to blogpost after all the processing"))

}

func GetPost(w http.ResponseWriter,r *http.Request){
	/*This handle function sends all the posts as response body 
	whenever the appropriate endpoint is hit via frontend*/
	setupCorsResponse(&w,r)
	json.NewEncoder(w).Encode(allPosts)

}