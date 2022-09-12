package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

var iv = []byte("tripleDE")  //8bytes of IV

type Request struct{
	Text string `json:"text"`
	Key string `json:"key"`
}

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/des/encrypt",TripleDESEncrypt).Methods("POST")
	r.HandleFunc("/des/decrypt",TripleDESDecrypt).Methods("POST")

	fmt.Println("Sever started at 8081")
	http.ListenAndServe(":8081",r)	
}

func createKey24(key string) []byte {
	keyHash := sha256.Sum256([]byte(key))
	//fmt.Printf("Total key is : %x", keyHash)
	//for 3DES, we need only 24bytes for key
	return keyHash[0:24]
	
}

func addPadding(data []byte,blockSize int) []byte{
	//PT should be multiple of 8 bytes so add padding accordingly
	noOfpaddingBytesToAdd := blockSize - len(data)%blockSize
	/*
	Either
	paddingText := bytes.Repeat([]byte{byte(noOfpaddingBytesToAdd)}, noOfpaddingBytesToAdd) 
	and then append it to data 
	or
	simply append the character 'X' to the data noOfPaddingBytesToAdd times.
	X is generally used due to it's rare possibilty of occuring in the message.
	*/
	sliceX := []byte{'X'} //1byte
	paddingText := bytes.Repeat(sliceX, noOfpaddingBytesToAdd)
	return append(data,paddingText...)

}

func TripleDESEncrypt(w http.ResponseWriter,r *http.Request){
	
	jsonData, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("err : ",err)
	}
	var receivedReq Request
	json.Unmarshal(jsonData, &receivedReq)

	tripleKey :=createKey24(receivedReq.Key)
	mytext := []byte(receivedReq.Text)
	block, _ := des.NewTripleDESCipher(tripleKey)
	fmt.Printf("%d bytes NewTripleDESCipher key with block size of %d bytes\n", len(tripleKey), block.BlockSize())

	//plaintext should be multiple of 8 bytes
	paddedPt := addPadding(mytext,block.BlockSize())
	encrypter := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(paddedPt))
	encrypter.CryptBlocks(ciphertext, paddedPt)
	
	fmt.Printf("%s encrypt to %x \n", receivedReq.Text, ciphertext)
	fmt.Println("CipherText in string format is : ",string(ciphertext))
	ctStr := hex.EncodeToString(ciphertext)

	w.Write([]byte(ctStr))
}


func TripleDESDecrypt(w http.ResponseWriter,r *http.Request){
	
	jsonData, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("err : ",err)
	}
	var receivedReq Request
	json.Unmarshal(jsonData, &receivedReq)
	data, err := hex.DecodeString(receivedReq.Text)
	//returns the bytes represented by the hexadecimal string receivedReq.Text.
	if err != nil {
   		 panic(err)
	}

	tripleKey :=createKey24(receivedReq.Key)
	block, _ := des.NewTripleDESCipher(tripleKey)
	decrypter := cipher.NewCBCDecrypter(block, iv)
	//allocating space to plainText
	cipherText := receivedReq.Text
    plainText := make([]byte, len(data))
	decrypter.CryptBlocks(plainText,data)

	fmt.Printf("%s decrypt to %x \n", cipherText, plainText)

	//remove the padded Xs considering that X was padded at the end while encryption
	res := strings.TrimRight(string(plainText), "X")

	w.Write([]byte(res))

}

