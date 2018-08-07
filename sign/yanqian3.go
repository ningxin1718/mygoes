//@Time  : 2018/3/9 14:59
//@Author: Greg Li
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var PrivKey *rsa.PrivateKey

type Message struct {
	Message string `json:"message"`
}

func (msg *Message) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(&msg)
}

type Signature struct {
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
	N         string `json:"N"`
	E         string `json:"E"`
}

func hash(msg string) []byte {
	sh := crypto.SHA1.New()
	sh.Write([]byte(msg))
	hash := sh.Sum(nil)
	return hash
}

func SignWithKey(msg Message) Signature {
	hash := hash(msg.Message)
	bytes, err := rsa.SignPKCS1v15(rand.Reader, PrivKey, crypto.SHA1, hash)
	if err != nil {
		panic(err)
	}
	signature := hex.EncodeToString(bytes)
	sig := Signature{
		hex.EncodeToString(hash),
		signature,
		PrivKey.PublicKey.N.String(),
		strconv.Itoa(PrivKey.PublicKey.E),
	}
	return sig
}

func sign(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/sign")
	var msg Message
	err := msg.Decode(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Signing: " + msg.Message)
	signature := SignWithKey(msg)
	js, err := json.Marshal(signature)
	fmt.Println(string(js))

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func LoadKeys() {
	// generate private key
	var err error
	PrivKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	fmt.Println("Loading Keys")
	LoadKeys()
	fmt.Println("Keys Loaded")
	http.HandleFunc("/sign", sign)
	http.ListenAndServe(":8080", nil)
}
