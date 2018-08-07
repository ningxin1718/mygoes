package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"fmt"
	"testing"
)

func main() {
	reader := rand.Reader
	bitSize := 2048
	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := &key.PublicKey
	savePEMKey("private.pem", key)
	savePublicPEMKey("public.pem", publicKey)
}



func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("private key err",err)
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err = pem.Encode(outFile, privateKey)
}

func savePublicPEMKey(fileName string, pubkey *rsa.PublicKey) {

	asn1Bytes, err :=x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		fmt.Println("public err: ",err)
	}
	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	pemfile, err := os.Create(fileName)
	defer pemfile.Close()
	err = pem.Encode(pemfile, pemkey)
}











