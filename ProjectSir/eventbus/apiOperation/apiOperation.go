package apioperation

import (
	"bytes"
	"encoding/json"
	// "encoding/json"
	// "encoding/json"
	"fmt"
	"io"
	// "log"
	"net/http"

	// "os"
	"github.com/gofrs/uuid"
)

// var posts =make(map[uuid.UUID]*Post)

// type Post struct{
// 	PostId uuid.UUID `json:"id"`
// 	Title string `json:"title"`
// }

//var queue[] Event


type Event struct{
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

func PostEvent(w http.ResponseWriter,r *http.Request){
	//iske request body me we have title type and id
	//create post event	
	//r.HandleFunc("/eventbus/event",apioperation.PostEvent).Methods("POST","OPTIONS") 

	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}

	jsonDataRbody, _ := io.ReadAll(r.Body)
	// jsonData2, _ := io.ReadAll(r.Body)

	fmt.Println("while sending r body to query service : ",string(jsonDataRbody))

	//send the json object to the eventbus
	fmt.Println("\nSending event from event bus to queryservice 4003 ",r.Body)
	resp, err := http.Post("http://localhost:4003/eventbus/event/listener", "application/json",bytes.NewBuffer(jsonDataRbody))
	if err != nil {
		fmt.Println("Inside eb :",err)
		//update in queue

	}
	if resp.Body!=nil && resp.StatusCode==200{
		jsonData, _ := io.ReadAll(resp.Body)
		fmt.Println("successfully posting event to qs.")
		fmt.Println(string(jsonData))
		//final ok to bp that everything is done bhaii..one full event cycle is done
		w.Write([]byte("OK from eb to bp or bc .After the whole event cycle is completed"))
	}
}

func AfterProcessEventFromQS(w http.ResponseWriter,r *http.Request){
	//	r.HandleFunc("/eventbus/getevent",apioperation.AfterProcessEventFromQS).Methods("POST","OPTIONS") 
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	jsonData, _ := io.ReadAll(r.Body)
	
	fmt.Println("\nGetting event in request body of EB \nSent event from QS to EB after updating QS db: ",string(jsonData))
    

	//depending on type-send a post request
	fmt.Println("now sent event back to blogpost again from EB")
	// var eventResponse Event;
	var SingleEventData Event
   	json.Unmarshal(jsonData,&SingleEventData)
	if SingleEventData.Type=="Post created"{
		resp, err := http.Post("http://localhost:4001/senteventafterprocess", "application/json",bytes.NewBuffer(jsonData))
		if resp.Body!=nil{
			//if ok from bp then send ok to qs
			//send 2nd ok response to QS
			jsonDataRes, _ := io.ReadAll(resp.Body)
			fmt.Println("LETS SEE : ",string(jsonDataRes))
			w.Write([]byte("Successfully got event from QS to EB"))

		}
		if err!=nil{
			fmt.Println("inside eb error : ",err)
		}
	}else{
		resp, err := http.Post("http://localhost:4002/senteventafterprocess", "application/json",bytes.NewBuffer(jsonData))
		if resp.Body!=nil{
			//if ok from bp then send ok to qs
			//send 2nd ok response to QS
			jsonDataRes, _ := io.ReadAll(resp.Body)
			fmt.Println("LETS SEE : ",string(jsonDataRes))
			w.Write([]byte("Successfully got event from QS to EB"))

		}
		if err!=nil{
			fmt.Println("inside eb error : ",err)
		}
	}


}






// func GetEvent(w http.ResponseWriter,r *http.Request){
// 	//	r.HandleFunc("/eventbus/event",apioperation.GetEvent).Methods("GET","OPTIONS")
// 	setupCorsResponse(&w,r)
// 	fmt.Println("Inside the get event of eventbus")
// 	json.NewEncoder(w).Encode(queue)
// }