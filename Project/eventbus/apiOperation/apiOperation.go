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

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func PostEvent(w http.ResponseWriter,r *http.Request){
	
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}

	jsonDataRbody, _ := io.ReadAll(r.Body)

	fmt.Println("while sending r body to query service 4003 : ",string(jsonDataRbody))

	resp, err := http.Post("http://localhost:4003/eventbus/event/listener", "application/json",bytes.NewBuffer(jsonDataRbody))
	if err != nil {
		fmt.Println("Inside EB :",err)
		var singleEvent model.Event
		json.Unmarshal(jsonDataRbody,&singleEvent)
		queue = append(queue, singleEvent)
	}
	if resp!=nil && resp.StatusCode==200{
		jsonData, _ := io.ReadAll(resp.Body)
		fmt.Println("successfully posting event from EB to QS.")
		fmt.Println(string(jsonData))
		//final ok to bp that everything is done bhaii..one full event cycle is done
		w.Write([]byte("OK from eb to bp or bc .After the whole event cycle is completed"))
	}
}

func AfterProcessEventFromQS(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)
	if(*r).Method=="OPTIONS"{
		return
	}
	jsonData, _ := io.ReadAll(r.Body)
	
	fmt.Println("\nGetting event in request body of EB \nSent event from QS to EB after updating QS db: ",string(jsonData))
    fmt.Println("Now send event back to blogpost again from EB")

	var SingleEvent model.Event
   	json.Unmarshal(jsonData,&SingleEvent)
	if SingleEvent.Type=="Post created"{
		resp, err := http.Post("http://localhost:4001/senteventafterprocess", "application/json",bytes.NewBuffer(jsonData))
		if resp!=nil{
			jsonDataRes, _ := io.ReadAll(resp.Body)
			fmt.Println("LETS SEE : ",string(jsonDataRes))
			w.Write([]byte("Successfully got event from QS to EB"))
		}
		if err!=nil{
			fmt.Println("inside eb error : ",err)
		}
	}else{
		resp, err := http.Post("http://localhost:4002/senteventafterprocess", "application/json",bytes.NewBuffer(jsonData))
		if resp!=nil{
			jsonDataRes, _ := io.ReadAll(resp.Body)
			fmt.Println("LETS SEE : ",string(jsonDataRes))
			w.Write([]byte("Successfully got event from QS to EB"))

		}
		if err!=nil{
			fmt.Println("inside eb error : ",err)
		}
	}


}


func GetEventQueue(w http.ResponseWriter,r *http.Request){
	setupCorsResponse(&w,r)

	defer func(){
		queue= nil
		fmt.Println("Queue is after nil statemnt : ",queue)

	}()

	fmt.Println("Inside the get event of eventbus.Send event queue to QS")
	json.NewEncoder(w).Encode(queue)
	
}