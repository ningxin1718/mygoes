//@Time  : 2018/4/3 15:43
//@Author: Greg Li
package main

import (
	"crypto/x509"
	"fmt"
	"encoding/pem"
	"io/ioutil"
)

func parseTcaEcdsa2() (rcaCert *x509.Certificate) {
	//解析tca证书
	tcaFile, err1 := ioutil.ReadFile("tca.crt")
	if err1 != nil {
		fmt.Println("ReadFile err:", err1)
		return
	}

	tcaBlock, _:= pem.Decode(tcaFile)
	if tcaBlock == nil {
		fmt.Println("ecaFile error")
		return
	}

	tcaCert, err2 := x509.ParseCertificate(tcaBlock.Bytes)
	if err2 != nil {
		fmt.Println("ParseCertificate err:", err2)
		return
	}
	return tcaCert
}

func main()  {
	tcaCert:=parseTcaEcdsa2()
	//验证签名
	//解析tx证书
	txFile, err := ioutil.ReadFile("0xd2a132139ca63447a7affc49143c17bf81948d54.tcert")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	txBlock, _:= pem.Decode(txFile)
	if txBlock == nil {
		fmt.Println("ecaFile error")
		return
	}

	txCert, err:= x509.ParseCertificate(txBlock.Bytes)
	if err != nil {
		fmt.Println("ParseCertificate err:", err)
		return
	}

	err = txCert.CheckSignatureFrom(tcaCert)
	fmt.Println("check txCert signature: ", err==nil)

	//解析地址
	addr2 := txCert.EmailAddresses
	fmt.Println("addr",addr2[0][:42])
	fmt.Println("level",addr2[0][43:44])
	fmt.Println("tag",addr2[0][44:45])

}