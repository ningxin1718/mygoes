package main

import (
	"encoding/hex"
	"fmt"
	"crypto/x509"
	"log"
	"io/ioutil"
	"encoding/pem"
)

func main()  {
	userCertString:=ReadUserCert()
	fmt.Println(userCertString)

	//x:=CheckUserCert(userCertString)
	x2:=parseRcaRsa2()
	fmt.Println(x2)
}

func parseRcaRsa2() bool{
	rcaFile, err := ioutil.ReadFile("rca.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
	}

	rcaBlock, _:= pem.Decode(rcaFile)
	if rcaBlock == nil {
		fmt.Println("ecaFile error")
	}

	rcaCert, err := x509.ParseCertificate(rcaBlock.Bytes)
	if err != nil {
		fmt.Println("ParseCertificate err:", err)
	}

	f, err := ioutil.ReadFile("user.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
	}

	rcaBlock2, _:= pem.Decode(f)
	if rcaBlock == nil {
		fmt.Println("ecaFile error")
	}

	userCert2, err := x509.ParseCertificate(rcaBlock2.Bytes)
	if err != nil {
		fmt.Println("ParseCertificate err:", err)
	}

	err = userCert2.CheckSignatureFrom(rcaCert)
	log.Println("check eCert signature: ", err==nil)
	return err==nil

}

func ReadUserCert() string {
	//BaseDir:=DefaultDataDir()
	f, err := ioutil.ReadFile("user.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
	}

	//Block, _:= pem.Decode(f)
	//if Block == nil {
	//	fmt.Println("ecaFile error")
	//}
	//
	//userCert:=pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: Block.Bytes})
	userCertString:=hex.EncodeToString(f)
	return userCertString
}

func parseRcaRsa() (rcaCert *x509.Certificate) {
	rcaFile, err := ioutil.ReadFile("rca.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	rcaBlock, _:= pem.Decode(rcaFile)
	if rcaBlock == nil {
		fmt.Println("ecaFile error")
		return
	}

	Cert, err := x509.ParseCertificate(rcaBlock.Bytes)
	if err != nil {
		fmt.Println("ParseCertificate err:", err)
		return
	}
	return Cert
}

func CheckUserCert(userCert string) bool{

	rcaCert := parseRcaRsa()

	//certToByte,err:=hexutil.Decode(userCert)
	certToByte,err:=hex.DecodeString(userCert)
	if err !=nil {
		fmt.Println("user's certificate format error")
		return false
	}
	fmt.Println(certToByte)
	userCert2, err := x509.ParseCertificate(certToByte)
	if err != nil {
		fmt.Println("ParseCertificate err:", err)
		return false
	}

	err = userCert2.CheckSignatureFrom(rcaCert)
	log.Println("check eCert signature: ", err==nil)
	return err==nil
}
