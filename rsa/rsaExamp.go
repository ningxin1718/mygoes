//@Time  : 2018/3/17 23:22
//@Author: Greg Li
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	signer, err := loadPrivateKey("private.pem")
	if err != nil {
		fmt.Println("signer is damaged: %v", err)
	}
	toSign := "date: Thu, 05 Jan 2012 21:31:40 GMT"
	signed, err := signer.Sign([]byte(toSign))
	if err != nil {
		fmt.Errorf("could not sign request: %v", err)
	}
	sig := base64.StdEncoding.EncodeToString(signed)
	//sig := hex.EncodeToString(signed)
	fmt.Printf("Encoded: %v\n", sig)




	parser, perr := loadPublicKey("public.pem")
	if perr != nil {
		fmt.Println("public could not sign request: %v", err)
	}



	err=parser.Unsign2([]byte(toSign),signed)
	fmt.Println("resule:",err)



	unsigned, err := parser.Unsign(signed);
	if err != nil {
		fmt.Errorf("could not sign request: %v", err)
	}

	fmt.Println("unsigned: ",unsigned)
	fmt.Printf("Decrypted: %v\n", base64.StdEncoding.EncodeToString(unsigned))
}


// loadPrivateKey loads an parses a PEM encoded private key file.
func loadPublicKey(path string) (Unsigner, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parsePublicKey(data)
}

// parsePublicKey parses a PEM encoded private key.
func parsePublicKey(pemBytes []byte) (Unsigner, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}
	fmt.Println(block)
	var rawkey interface{}
	switch block.Type {
	case "PUBLIC KEY":
		rsaa, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		rawkey = rsaa
	default:
		fmt.Println("JKFjlasdkjflasdfasd")
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
	return newUnsignerFromKey(rawkey)
}


// loadPrivateKey loads an parses a PEM encoded private key file.
func loadPrivateKey(path string) (Signer, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return parsePrivateKey(data)
}

// parsePublicKey parses a PEM encoded private key.
func parsePrivateKey(pemBytes []byte) (Signer, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}
	var rawkey interface{}
	switch block.Type {
	case "PRIVATE KEY":
		rsaa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsaa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
	return newSignerFromKey(rawkey)
}

// A Signer is can create signatures that verify against a public key.
type Signer interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Sign(data []byte) ([]byte, error)
}

// A Signer is can create signatures that verify against a public key.
type Unsigner interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Unsign(data []byte) ([]byte, error)
	Unsign2(message []byte, sig []byte) error
}

func newSignerFromKey(k interface{}) (Signer, error) {
	var sshKey Signer
	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

func newUnsignerFromKey(k interface{}) (Unsigner, error) {
	var sshKey Unsigner
	switch t := k.(type) {
	case *rsa.PublicKey:
		sshKey = &rsaPublicKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

type rsaPublicKey struct {
	*rsa.PublicKey
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

// Sign signs data with rsa-sha256
func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d)
}

// Unsign encrypts data with rsa-sha256
func (r *rsaPublicKey) Unsign(message []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, r.PublicKey, message)
}

// Unsign verifies the message using a rsa-sha256 signature
func (r *rsaPublicKey) Unsign2(message []byte, sig []byte) error {
	h := sha256.New()
	h.Write(message)
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, d, sig)
}
