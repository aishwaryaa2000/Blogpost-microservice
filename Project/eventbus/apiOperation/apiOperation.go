package apioperation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"eventbus/model"
)

var queue =make([]model.Event, 0)
//This queue is used to store the events when query service goes down so that these events can be served later

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	//this is used to avoid cors error
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func PostEvent(w http.ResponseWriter,r *http.Request){
	/*
	This handle func is used to send the events recieved from blogpost/blogcomment to queryservice
	*/
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}

	jsonDataEvent, _ := io.ReadAll(r.Body)
	fmt.Println("while sending r body to query service 4003 : ",string(jsonDataEvent))

	resp, err := http.Post("http://query-serv:4003/eventbus/event/listener", "application/json",bytes.NewBuffer(jsonDataEvent))
	if err != nil {
		//When queryService is down,then store the events into queue so that it can be served later when QS is up
		fmt.Println("Inside EB :",err)
		var singleEvent model.Event
		json.Unmarshal(jsonDataEvent,&singleEvent)
		queue = append(queue, singleEvent)
	}
	if resp!=nil && resp.StatusCode==200{
		jsonData, _ := io.ReadAll(resp.Body)
		fmt.Println("After successfully posting event from eventbus to QueryService.")
		fmt.Println("Response in eventbus from queryService : ",string(jsonData))
		w.Write([]byte("OK from eventbus to blogpost or blogcomment .After the whole event cycle is completed"))
	}
}

func RedirectProcessedEvent(w http.ResponseWriter,r *http.Request){
	/*This handle function is to redirect the processed event 
	from queryService to blogpost or blogcomment according to the event type*/
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	jsonData, _ := io.ReadAll(r.Body)
	fmt.Println("\nGetting event in request body of EventBus \nThis is the event sent from QS to EB after updating QS db: ",string(jsonData))

	var SingleEvent model.Event
   	json.Unmarshal(jsonData,&SingleEvent)

	//redirecting event acc to event type
	if SingleEvent.Type=="Post created"{
		resp, err := http.Post("http://posts-serv:4001/processedevent", "application/json",bytes.NewBuffer(jsonData))
		if resp!=nil{
			jsonDataRes, _ := io.ReadAll(resp.Body)
			fmt.Println("Response after sending processed event from EB to BP : ",string(jsonDataRes))
			w.Write([]byte("Successfully got processed event from QS to EB"))
		}
		if err!=nil{
			fmt.Println("Inside eb error from BP : ",err)
		}
	}else{
		//if type of the event is comment created
		resp, err := http.Post("http://comment-serv:4002/processedevent", "application/json",bytes.NewBuffer(jsonData))
		if resp!=nil{
			jsonDataRes, _ := io.ReadAll(resp.Body)
			fmt.Println("Response after sending processed event from EB to BC  : ",string(jsonDataRes))
			w.Write([]byte("Successfully got processed event from QS to EB"))

		}
		if err!=nil{
			fmt.Println("Inside eb error from BC : ",err)
		}
	}

}


func SendEventQueue(w http.ResponseWriter,r *http.Request){
	//This handle func is used to send the queue to QS

	setupCorsResponse(&w,r)

	/*After giving the queue to the QS in the response body
	  empty the queue as the events of the queue will be 
	  processed by the QS
	*/
	defer func(){
		queue= nil
		fmt.Println("Queue is after nil statemnt : ",queue)

	}()
	
	json.NewEncoder(w).Encode(queue)
	
}