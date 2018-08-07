//@Time  : 2018/5/7 14:47
//@Author: Greg Li
package main

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"encoding/pem"
)

func main() {

	//2.解析test证书///////////////////////////////////////////////////
	tFile, err := ioutil.ReadFile("0x672f7c84c334765e7828ba6d51570709a8322ae4.tcert")
	if err != nil {
		fmt.Println("ReadFile err:", err)
	}

	tBlock, _:= pem.Decode(tFile)
	if tBlock == nil {
		fmt.Println("tFile error")
		return
	}

	tCert, err := x509.ParseCertificate(tBlock.Bytes)
	if err != nil {
		fmt.Println("ParseCertificate err:", err)
	}

	fmt.Println(tCert.EmailAddresses)
}
