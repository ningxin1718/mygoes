//@Time  : 2018/3/23 11:33
//@Author: Greg Li
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"crypto/ecdsa"
)

func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
	pubBytes := crypto.FromECDSAPub(&p)
	fmt.Println("pubBytes:",pubBytes)
	fmt.Println(len(pubBytes))

	pub:=crypto.Keccak256(pubBytes[:])
	fmt.Println("pub---->",pub)
	fmt.Println("len(pub):",len(pub))

	x:=common.BytesToAddress(crypto.Keccak256(pubBytes[1:]))
	fmt.Println(x)
	fmt.Println(len(x))

	return common.BytesToAddress(crypto.Keccak256(pubBytes[1:])[12:])
}

func main()  {
	key,_:=crypto.HexToECDSA("03f6fc089d1f27943506638f0b44c804ca2d8feee13f488116d7b3fd82e831c7")
	fmt.Println("key===========>",key)

	// 得到地址：42个字符
	fmt.Println("key.PublicKey:",key.PublicKey)





	address := PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println(address)
}
