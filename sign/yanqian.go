//@Time  : 2018/3/7 14:26
//@Author: Greg Li
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"encoding/pem"
	"errors"
)

func sign(hashed []byte) ([]byte,error){
	block, _ := pem.Decode(PrivateKeyData)
	if block == nil {
		fmt.Println("public key error")
		return nil,nil
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil,nil
	}

	//签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
	if err !=nil {
		fmt.Println("signature err",err)
	}
	return signature,err
}

func verifySign(signature []byte,hashed []byte) ([]byte, error)  {
	block, _:= pem.Decode(PublicKeyData)
	if block == nil {
		fmt.Println("block nil")
		return nil,errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil,err
	}

	pub := pubInterface.(*rsa.PublicKey) //pub:公钥对象
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, signature)
	if err!=nil {
		fmt.Println("Verify error:",err)
		return nil,err
	}
	fmt.Println("success")
	return nil,err
}

func main() {
	strdata := "asdf"//待签名数据
	h2 := sha256.New()
	h2.Write([]byte(strdata))
	hashed := h2.Sum(nil)

	signature,err := sign(hashed)
	if err !=nil {
		fmt.Println(err)
	}

	verifySign(signature,hashed)
}


var PrivateKeyData = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)

var PublicKeyData = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)