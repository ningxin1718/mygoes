//@Time  : 2018/3/7 15:06
//@Author: Greg Li
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
)

func main() {
	block2, _ := pem.Decode(PrivateKeyData)//PiravteKeyData为私钥文件的字节数组
	if block2 == nil {
		fmt.Println("block空")
		return
	}
	//priv即私钥对象,block2.Bytes是私钥的字节流
	priv, err := x509.ParsePKCS1PrivateKey(block2.Bytes)
	if err != nil {
		fmt.Println("无法还原私钥")
		return
	}

	strdata := "asdf"//待签名数据
	h2 := sha256.New()
	h2.Write([]byte(strdata))
	hashed := h2.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)//签名
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hashed,signature)



	//下面开始验签，PublicKeyData为公钥文件的字节数组
	block, _ := pem.Decode(PublicKeyData)
	if block == nil {
		fmt.Println("block nil")
		return
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("还原公钥错误",err)
		return
	}

	pub := pubInterface.(*rsa.PublicKey)//pub:公钥对象
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, signature)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
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