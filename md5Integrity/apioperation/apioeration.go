package apioperation

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct{
	Text string `json:"text"`
}


func createHashMsg(msg string) string {
	hasher := md5.New()
	hasher.Write([]byte(msg))
	return hex.EncodeToString(hasher.Sum([]byte(msg)))
	//Sum appends the current hash to b and returns the resulting slice.
}


func createHash(msg string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(msg))
	return (hasher.Sum([]byte(nil)))
}


func SendMsgHash(w http.ResponseWriter,r *http.Request){

	jsonData, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("err : ",err)
	}
	var receivedReq Request
	json.Unmarshal(jsonData, &receivedReq)
	msgHash := createHashMsg(receivedReq.Text)
	w.Write([]byte(msgHash))

}

func CheckMsgHash(w http.ResponseWriter,r *http.Request){
	jsonData, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("err : ",err)
	}

	var receivedReq Request
	json.Unmarshal(jsonData, &receivedReq)
	receivedMsgWithHash, _ := hex.DecodeString(receivedReq.Text)

	receivedHash := receivedMsgWithHash[len(receivedMsgWithHash)-16:]
	receivedMsg := receivedMsgWithHash[:len(receivedMsgWithHash)-16]
	generatedHash := createHash(string(receivedMsg))

	if(string(generatedHash)==string(receivedHash)){
		w.Write([]byte("Data integrity maintained"))

	}else{
		w.Write([]byte("Data integrity lost"))
	}


}