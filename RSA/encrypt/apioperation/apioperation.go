package apioperation

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct{
	Text string `json:"text"`
}

func encrypt(msg string, publickey *rsa.PublicKey) string {
	cipher, _ := rsa.EncryptOAEP(sha1.New(), rand.Reader, publickey, []byte(msg), nil) //encryption
	return string(cipher)
}


func SendMsg(w http.ResponseWriter,r *http.Request){

	jsonData, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("err : ",err)
	}
	var receivedReq Request

	json.Unmarshal(jsonData, &receivedReq)
	resp,err := http.Get("http://localhost:8081/rsa/publickey")
	if err!=nil{
		fmt.Println("Error while trying to get public key : ",err)
	}

	var PublicKeyReceived rsa.PublicKey
	err = json.NewDecoder(resp.Body).Decode(&PublicKeyReceived)
	if err != nil {
		fmt.Println("err while decoding public key: ",err)
	}
	fmt.Println("Public Key received : ",PublicKeyReceived)
	cipherText := encrypt(receivedReq.Text,&PublicKeyReceived)
	cipher := hex.EncodeToString([]byte(cipherText))

	w.Write([]byte(cipher))

}