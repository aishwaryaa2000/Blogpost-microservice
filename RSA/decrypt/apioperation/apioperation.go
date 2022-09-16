package apioperation


import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
)

type Request struct{
	Text string `json:"text"`
}

var privatekey, _ = rsa.GenerateKey(rand.Reader, 2048)
var publickey = privatekey.PublicKey

func SendPublicKey(w http.ResponseWriter,r *http.Request){

	err := json.NewEncoder(w).Encode(publickey)
	if err!=nil{
		fmt.Println("Error while sending publickey : ",err )
	}

}

func decrypt(msg string, privatekey *rsa.PrivateKey) []byte {
	plaintext, _ := rsa.DecryptOAEP(sha1.New(), rand.Reader, privatekey, []byte(msg), nil)
	fmt.Println("PT inside decrypt() : ",string(plaintext)) // decryption
	return plaintext
}


func GeneratePlaintext(w http.ResponseWriter,r *http.Request){
	var receivedReq Request
	err := json.NewDecoder(r.Body).Decode(&receivedReq)
	if err!=nil{
		fmt.Println("Error while receiving CT : ",err)
	}
	cipher, _ := hex.DecodeString(receivedReq.Text)

	plainText := decrypt(string(cipher), privatekey)

	w.Write(plainText)

}