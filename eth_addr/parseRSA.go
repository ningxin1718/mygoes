//@Time  : 2018/3/14 11:34
//@Author: Greg Li
package main

import (
	"crypto/x509"
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	//1.解析rca证书
	rcaFile, err1 := ioutil.ReadFile("rootca.cer")
	if err1 != nil {
		fmt.Println("ReadFile err:", err1)
		return
	}

	//rcaBlock, _:= pem.Decode(rcaFile)
	//if rcaBlock == nil {
	//	fmt.Println("rcaFile error")
	//	return
	//}

	rcaCert, err := x509.ParseCertificate(rcaFile)
	if err != nil {
		fmt.Println("ParseCertificate err:", err)
		return
	}

    //2.解析test证书///////////////////////////////////////////////////
	tFile, err := ioutil.ReadFile("test1.cer")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	//tBlock, _:= pem.Decode(tFile)
	//if tBlock == nil {
	//	fmt.Println("tFile error")
	//	return
	//}

	tCert, err := x509.ParseCertificate(tFile)
	if err != nil {
		fmt.Println("ParseCertificate err:", err)
		return
	}

	//查看证书内容////////////////////////////////////////////////
	fmt.Println(tCert.EmailAddresses)
	fmt.Println(tCert.Subject)
	fmt.Println(tCert.Subject.CommonName)
	fmt.Println(tCert.Subject.Country)
	fmt.Println(tCert.Subject.Organization)

	fmt.Println(rcaCert.EmailAddresses)
	fmt.Println(rcaCert.Subject)
	fmt.Println(rcaCert.Subject.CommonName)
	fmt.Println(rcaCert.Subject.Country)
	fmt.Println(rcaCert.Subject.Organization)

	//check //////////////////////////////////////////////////////////
	err = tCert.CheckSignatureFrom(rcaCert)
	log.Println("check tCert signature: ", err==nil)
}