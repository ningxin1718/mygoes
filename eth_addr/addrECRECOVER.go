//@Time  : 2018/3/23 11:33
//@Author: Greg Li
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/sha3"

)

func main()  {
	key,_:=crypto.HexToECDSA("03f6fc089d1f27943506638f0b44c804ca2d8feee13f488116d7b3fd82e831c7")
	fmt.Println("key===========>",key)

	//公钥
	pub:=common.ToHex(crypto.FromECDSAPub(&key.PublicKey))
	pubBytes := crypto.FromECDSAPub(&key.PublicKey)

	fmt.Println("pubBytes::",pubBytes)

	c:=crypto.Keccak256(pubBytes[1:])
	d:=crypto.Keccak256(pubBytes[1:])[12:]
	fmt.Println("pub------>",pub)
	fmt.Println("Hex(c):::::",Hex(c))
	fmt.Println("Hex(d):::::",Hex(d))


	// 得到地址：42个字符
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println(address)


	//签名
	msg := crypto.Keccak256([]byte(address))
	msg2:=hexutil.Encode(msg)
	fmt.Println("msg--->",msg2)
	msg3,_:=hexutil.Decode(msg2)

	priv2:=math.PaddedBigBytes(key.D, 32)
	signature,err2 := secp256k1.Sign(msg,priv2)
	fmt.Println("signature:::",signature)
	fmt.Println(len(signature))  //65
	r := signature[0:32]
	s:=signature[32:64]
	fmt.Println(r)
	fmt.Println(s)

	sig2:=hexutil.Encode(signature)
	fmt.Println("sig--->",sig2)
	sig3,_:=hexutil.Decode(sig2)

	//signature,err2:=crypto.Sign(msg,priv)
	if err2 !=nil {
		fmt.Println(err2)
	}

	err3:=secp256k1.VerifySignature(pubBytes, msg3, sig3[:64])
	fmt.Println(err3)

	recoverpub,_:=crypto.Ecrecover(msg3,sig3)
	fmt.Println("recover pub:",recoverpub)
}

func Hex(a []byte) string {
	unchecksummed := hex.EncodeToString(a[:])
	sha := sha3.NewKeccak256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return "0x" + string(result)
}