package main

import (
	"strings"
	"fmt"
	"crypto/ecdsa"
	"github.com/usechain/go-usechain/common"
	"github.com/usechain/go-usechain/crypto"
	"github.com/btcsuite/btcd/btcec"
	"encoding/hex"
)

func main (){
	ASset:="02eb0165fa960d243932b2fd8ba55d9c0ac07ec17b3ba1dbe5f40cb77aac1f920c03e60f33281bd0cbdf4865a8ca4eee01322faf0be944805a99ac5ead458582ea68,02f2026857f3708d66acb5fdfd9079644fc58ccf0008c427898585478af19ea475027e70238df299395f1b3783441430b66693fdd9362b03bb2f4d09400804bd3f1b"

	ASslice := strings.Split(ASset,",")
	publicKeyset := make([]string, 0)
	for _, AS := range ASslice {
		ASbyte,_:= hex.DecodeString(AS)
		pk1,_,err := GeneratePKPairFromABaddress(ASbyte)
		fmt.Println(pk1)
		if err !=nil {
		}
		pub:=common.ToHex(crypto.FromECDSAPub(pk1))
		fmt.Println("pub:::",pub)
		publicKeyset = append(publicKeyset, pub)
	}
	publickeys := strings.Join(publicKeyset, ",")
	fmt.Println("pub=========================================",publickeys)
}


func GeneratePKPairFromABaddress(w []byte) (*ecdsa.PublicKey, *ecdsa.PublicKey, error) {
	if len(w) != common.ABaddressLength {
		fmt.Println(len(w))
		return nil, nil, nil
	}

	tmp := make([]byte, 33)
	copy(tmp[:], w[:33])
	curve := btcec.S256()
	PK1, err := btcec.ParsePubKey(tmp, curve)
	if err != nil {
		//这里遇到错误：tmp全是0：invalid magic in compressed pubkey string: 0
		return nil, nil, err
	}

	copy(tmp[:], w[33:])
	PK2, err := btcec.ParsePubKey(tmp, curve)
	if err != nil {
		return nil, nil, err
	}

	PK11:=(*ecdsa.PublicKey)(PK1)
	PK22:=(*ecdsa.PublicKey)(PK2)
	return PK11, PK22, nil
}