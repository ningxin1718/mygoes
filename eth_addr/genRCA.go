//@Time  : 2018/3/13 11:43
//@Author: Greg Li
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"math/big"
	"os"
	"io/ioutil"
	"time"
	"encoding/pem"
	"github.com/astaxie/beego/logs"
	"crypto/ecdsa"
)

func main(){
	//GenRCA("292409083@qq.com",true,"rca2xu.crt")
	GenTCA(true,"xutonghua20180508",1)
}

func GenRCA(emailAddress string,isCA bool,caName string) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err !=nil {
		fmt.Println(err)
	}
	pub := publicKey(key)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Country: []string{"CN"},
			Organization: []string{"UseChain"},
		},
		NotBefore: time.Now(),
		NotAfter: time.Now().AddDate(10,0,0),

		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}
	template.DNSNames = append(template.DNSNames, "localhost")
	template.EmailAddresses = append(template.EmailAddresses, emailAddress)

	if isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}
	rcaCert, err := x509.CreateCertificate(rand.Reader, &template, &template, pub, key)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}
	SaveBinaryCA(caName,rcaCert)
	savePrivkey("key2xu.key",key)
}

func SaveCA(caName string,cert []byte ){
	certOut, err := os.Create(caName)
	if err != nil {
		log.Fatalf("failed to open for writing: %s", err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
	logs.Debug("write ecert to", certOut.Name())
	certOut.Close()
}

func SaveBinaryCA(caName string,cert []byte ){
	certOut, err := os.Create(caName)
	if err != nil {
		log.Fatalf("failed to open for writing: %s", err)
	}
	//pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
	//logs.Debug("write ecert to", cert.Name())
	_, writeErr := certOut.Write(cert)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
	certOut.Close()
}

func GenTCA(isCA bool,caName string,lev int) {

	var err error
	var rcaCert *x509.Certificate
	var rcaPrivKey interface{}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	rcaCert,rcaPrivKey = parseRsa()

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Country: []string{"CN"},
			Organization: []string{"UseChain"},
		},
		NotBefore: time.Now(),
		NotAfter: time.Now().AddDate(10,0,0),

		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}
	template.DNSNames = append(template.DNSNames, "localhost")
	addr1:=fmt.Sprintf("0x00000000000000000000000000000000000000b1@%v.com",lev)
	fmt.Println(lev)
	template.EmailAddresses = append(template.EmailAddresses, addr1)
	if isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	pub := publicKey(priv)

	tcaCert, err := x509.CreateCertificate(rand.Reader, &template, rcaCert, pub, rcaPrivKey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	SaveBinaryCA(caName,tcaCert)

	//savePrivkey(privName,priv)
	//if pubName=="" {
	//
	//} else {
	//	savePubPem(pubName,pub)
	//}

	tcaCert2, _ := x509.ParseCertificate(tcaCert)
	err = tcaCert2.CheckSignatureFrom(rcaCert)
	log.Println("check tcaCert signature: ", err==nil)
}

func parseRsa() (rcaCert *x509.Certificate,rcaPrivKey *rsa.PrivateKey) {
	//解析rca证书
	rcaFile, err1 := ioutil.ReadFile("rca2xu.crt")
	if err1 != nil {
		fmt.Println("ReadFile err:", err1)
		return
	}

	//rcaBlock, _:= pem.Decode(rcaFile)
	//if rcaBlock == nil {
	//	fmt.Println("ecaFile error")
	//	return
	//}

	rcaCert, err2 := x509.ParseCertificate(rcaFile)
	if err2 != nil {
		fmt.Println("ParseCertificate err:", err2)
		return
	}

	//解析rca私钥
	rcaPriv, err3 := ioutil.ReadFile("key2xu.key")
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	rcaKeyBlock, _ := pem.Decode(rcaPriv)
	if rcaKeyBlock == nil {
		fmt.Println("ecaKeyBlock nil error")
		return
	}

	rcaPrivKey, parseErr := x509.ParsePKCS1PrivateKey(rcaKeyBlock.Bytes)
	if parseErr != nil {
		fmt.Println(parseErr)
		return
	}
	return rcaCert,rcaPrivKey
}

func savePrivkey(privName string,priv interface{} )  {
	keyOut, err := os.OpenFile(privName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Print("failed to open key.pem for writing:", err)
		return
	}
	pem.Encode(keyOut, pemBlockForKey(priv))
	keyOut.Close()
	log.Println("write to", keyOut.Name())
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}