//@Time  : 2018/3/8 9:20
//@Author: Greg Li
package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"crypto/rsa"
	"crypto/rand"
	"math/big"
	"io/ioutil"
	"time"
	"log"
	"fmt"
	"encoding/pem"
	"os"
)

func main()  {
	//解析eca证书
	ecaFile, err := ioutil.ReadFile("eca.pem")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	ecaBlock, _:= pem.Decode(ecaFile)
	if ecaBlock == nil {
		fmt.Println("ecaFile error")
		return
	}

	ecaCert, err2 := x509.ParseCertificate(ecaBlock.Bytes)
	if err2 != nil {
		fmt.Println("ParseCertificate err:", err)
		return
	}

	//解析eca私钥
	ecaPriv, err := ioutil.ReadFile("eca.key")
	if err != nil {
		fmt.Println(err)
		return
	}
	ecaKeyBlock, _ := pem.Decode(ecaPriv)
	ecaPrivKey, err := x509.ParsePKCS1PrivateKey(ecaKeyBlock.Bytes)
	if err != nil {
		return
	}

	//用ECA生成的三级用户证书
	userEcaCert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Country: []string{"CN"},
			Organization: []string{"UsechainUserECA"},
			OrganizationalUnit: []string{"UserECA"},
		},
		NotBefore: time.Now(),
		NotAfter: time.Now().AddDate(10,0,0),
		SubjectKeyId: []byte{1,2,3,4,6},
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage: x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign,
	}

	//生成公钥私钥对
	userEcaPriv, err:= rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	pub := &userEcaPriv.PublicKey
	userEcaCert_b, err:= x509.CreateCertificate(rand.Reader, userEcaCert,ecaCert,pub, ecaPrivKey)
	if err != nil {
		fmt.Println("CreateCertificate err:",err)
		return
	}

	//保存用户证书到文件
	userEcaCert_f, err := os.Create("userEca.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(userEcaCert_f, &pem.Block{Type: "CERTIFICATE", Bytes: userEcaCert_b})
	userEcaCert_f.Close()
	log.Println("write to", userEcaCert_f)


	//保存userEcaPriv私钥到文件
	userKeyFile, err := os.Create("userEca.key")
	if err != nil {
		panic(err)
	}
	userEcaPrivBuf := x509.MarshalPKCS1PrivateKey(userEcaPriv)
	userEcaPrivKeyPem := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: userEcaPrivBuf,
	}
	pem.Encode(userKeyFile, userEcaPrivKeyPem)
	log.Println("write to", userKeyFile.Name())
	userKeyFile.Close()


	//ecaParse, _ := x509.ParseCertificate(ecaCert_b)
	userEcaCert_c, _ := x509.ParseCertificate(userEcaCert_b)
	err1 := userEcaCert_c.CheckSignatureFrom(ecaCert)
	log.Println("check signature", err1 == nil)
}
