//@Time  : 2018/3/23 11:33
//@Author: Greg Li
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"math/big"
)

var (
	secp256k1_N, _  = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1_halfN = new(big.Int).Div(secp256k1_N, big.NewInt(2))
)


func main()  {
	fmt.Println("secp256k1_N::",secp256k1_N)
	fmt.Println("secp256k1_halfN::",secp256k1_halfN)

	decodePrivate,_:=hex.DecodeString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
	fmt.Println("decodePrivate::",decodePrivate)

	key,err:=crypto.HexToECDSA("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println("key===========>",key)


	//公钥
	//pub:=common.ToHex(crypto.FromECDSAPub(&key.PublicKey))

	pub:=&key.PublicKey
	pub1X := hexutil.Encode(common.LeftPadBytes(pub.X.Bytes(), 32))
	pub1Y := hexutil.Encode(common.LeftPadBytes(pub.Y.Bytes(), 32))

	fmt.Println(pub1X)
	fmt.Println(pub1Y)


	pubBytes := crypto.FromECDSAPub(&key.PublicKey)
	c:=crypto.Keccak256(pubBytes[1:])
	d:=crypto.Keccak256(pubBytes[1:])[12:]
	fmt.Println("pub------>",pub)
	fmt.Println("pub------>",&key.PublicKey)
	fmt.Println(Hex(c))
	fmt.Println(Hex(d))

	// 得到地址：42个字符
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println(address)

	//公钥hash
	//testpubkey ,_ := hexutil.Decode(pub)
	//pub2:=crypto.FromECDSAPub(&key.PublicKey)
	//fmt.Println(pub2)
	//fmt.Println("testpubkey----======》",testpubkey)
	//pubaddr:=crypto.Keccak256(testpubkey[1:])[12:]
	//fmt.Println(Hex(pubaddr))

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