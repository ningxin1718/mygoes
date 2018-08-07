package main

import (
	"crypto/x509"
	"io/ioutil"
	"fmt"
	"encoding/pem"
	"encoding/hex"
	"crypto/sha256"
	"crypto/rsa"
	"crypto"
)

type rsaPublicKey struct {
	*rsa.PublicKey
}


func main() {
	//解析rca证书
	rcaFile, err1 := ioutil.ReadFile("user.crt")
	if err1 != nil {
		fmt.Println("ReadFile err:", err1)
		return
	}

	rcaBlock, _:= pem.Decode(rcaFile)
	if rcaBlock == nil {
		fmt.Println("ecaFile error")
		return
	}

	rcaCert, err2 := x509.ParseCertificate(rcaBlock.Bytes)
	if err2 != nil {
		fmt.Println("ParseCertificate err:", err2)
	}


	fmt.Println(rcaCert.PublicKey)
	fmt.Println(rcaCert.SignatureAlgorithm)
	message:="21a7325a75492db1ee86c1d2d22984b83f6082e5"
	sig:="393ec9e9e2d48daa5ff3b14b7441e2c95009d5a913420bb0b7c8a059100295674a461394f174e1366d5bd604ff123bde4e05cedea033ac5c1f28b070f1f2abc45b7463364f8a4aacee7c9caa8661ee28fac4065902b833af35bfcd0dfce3c7695658ba28a89bfad51bfe30a620a3cd9112aa982d9d62c7f9c82287fb357d9fb8127c063511e25108ff50022c59efcb92243f864f4d24ab58d9e15a7da5e0d26dc85230489c5c31753d343bd268e1a45e41798e760bedf8fa513a6545d7da01488e23418c206555a70f3fa56b3d0f9f01c2a0d73acb2a513ece2f1dff85d9d5bf508e679f194546c613487aa284b6e311abb22f3e9267caf7dce84eb5c71a20fb"

	k:=rcaCert.PublicKey
	switch t := k.(type) {
	case *rsa.PublicKey:
		sshKey := &rsaPublicKey{t}
		err:=RSA_Verify(message,sig,sshKey.PublicKey)
		fmt.Println("a::::::::::::::::::::::::::::::",err)
	default:
		fmt.Println("adfasdf")
	}

}


func RSA_Verify(message string,sig string,pub *rsa.PublicKey) bool {
	signed,err:=hex.DecodeString(sig)
	if err != nil {
		fmt.Errorf("could not sign request: %v", err)
	}

	h := sha256.New()
	h.Write([]byte(message))
	d := h.Sum(nil)

	rsa.VerifyPKCS1v15(pub, crypto.SHA256, d, signed)
	return err==nil
}


