//@Time  : 2018/3/23 11:33
//@Author: Greg Li
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/usechain/go-usechain/crypto"
	"github.com/usechain/go-usechain/crypto/sha3"
	"github.com/usechain/go-usechain/common"
	"go-wanchain-feature-scan/common/hexutil"
)

func main()  {
	////// 创建账户
	//key, err := crypto.GenerateKey()
	//if err !=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("key_______________",key)
	//// 私钥:64个字符
	//privateKey := hex.EncodeToString(key.D.Bytes())
	//fmt.Println("private----->",privateKey)
	//decodePrivate,_:=hex.DecodeString(privateKey)
	//fmt.Println(decodePrivate)

	key,_:=crypto.HexToECDSA("e3dc14a49229f85f90c6156ca5fdfcd91e1e131700b60b0eb51cc1662af04713")
	fmt.Println("key===========>",&key)

	//key:=key2.(*ecdsa.PrivateKey)

	//公钥
	fmt.Println(&key.PublicKey)
	pub:=common.ToHex(crypto.FromECDSAPub(&key.PublicKey))
	fmt.Println("pub------>",pub)


	pubBytes := crypto.FromECDSAPub(&key.PublicKey)
	pub22:=hexutil.Encode(pubBytes)
	fmt.Println("pub22::::::::",pub22)


	c:=crypto.Keccak256(pubBytes[1:])
	d:=crypto.Keccak256(pubBytes[1:])[12:]
	fmt.Println("pub------>",&key.PublicKey)
	fmt.Println(Hex(c))
	fmt.Println(Hex(d))


	// 得到地址：42个字符
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println(address)
	fmt.Println(crypto.PubkeyToAddress(key.PublicKey))

	pubstring:="0x04f2026857f3708d66acb5fdfd9079644fc58ccf0008c427898585478af19ea4752125582144579c8be99f8f5bc2eb39dc5ea6fd9917ed558a489d7894780c0b18"
	pubstringTObyte,_:=hexutil.Decode(pubstring)
	pubxxx:=crypto.ToECDSAPub(pubstringTObyte)
	fmt.Println("pubxxx：：：：",pubxxx)

	pubaddr:=crypto.Keccak256(pubstringTObyte[1:])[12:]
	fmt.Println(pubaddr)
	fmt.Println(Hex(pubaddr))

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
