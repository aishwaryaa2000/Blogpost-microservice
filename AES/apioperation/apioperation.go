package apioperation

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type Request struct{
	Text string `json:"text"`
	Key string `json:"key"`
}


func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte,passphrase string) []byte{
	//First passphrase has to be hashed in order to get the key
	//and then key is passed to AES
	block,_ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm,_ := cipher.NewGCM(block)
	nonce := make([]byte,gcm.NonceSize())
	io.ReadFull(rand.Reader,nonce)
	cipherText := gcm.Seal(nonce,nonce,data,nil)
	return cipherText

}

func decrypt(data []byte,passphrase string) []byte{
	key := []byte(createHash(passphrase))
	block,_ := aes.NewCipher(key)
	gcm,_ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce:= data[:nonceSize]
	cipherText := data[nonceSize:]
	plainText,_ := gcm.Open(nil,nonce,cipherText,nil)
	return plainText 
}



func TripleAESEncrypt(w http.ResponseWriter,r *http.Request){
	
	jsonData, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("err : ",err)
	}
	var receivedReq Request
	json.Unmarshal(jsonData, &receivedReq)

	aesKey :=createHash(receivedReq.Key)
	mytext := []byte(receivedReq.Text)
	cipherText:=encrypt(mytext,aesKey)
	ctStr := hex.EncodeToString(cipherText)
	w.Write([]byte(ctStr))
}


func TripleAESDecrypt(w http.ResponseWriter,r *http.Request){
	
	jsonData, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("err : ",err)
	}
	var receivedReq Request
	json.Unmarshal(jsonData, &receivedReq)
	cipherText, err := hex.DecodeString(receivedReq.Text)
	//returns the bytes represented by the hexadecimal string receivedReq.Text.
	if err != nil {
   		 panic(err)
	}
	aesKey :=createHash(receivedReq.Key)
	ptStr := decrypt(cipherText,aesKey)

	w.Write([]byte(ptStr))

}

