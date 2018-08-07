package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"encoding/binary"
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	HashLength    = 32
)

//type Hash [HashLength]byte
//contractAddress := ""

type storjFlag struct {
	index 		uint
	parameter 	uint
	dimension	uint
	returns		string
}

var (
	OneTimeAddrConfirmedLenIndex = storjFlag{
		index:3,
		parameter:0,
	}
	confirmedMainAddressLenIndex = storjFlag{
		index:4,
		parameter:0,
	}
	confirmedSubAddressLenIndex = storjFlag{
		index:5,
		parameter:0,
	}
	isCommitteeIndex = storjFlag{
		index:6,
		parameter:2,
	}
	CMMTTEEs = storjFlag{
		index:7,
		parameter:15,
	}
	CommitteePublicKey = storjFlag{
		index:8,
		parameter:0,
	}
	CertToAddress = storjFlag{
		index:9,
		parameter:2,
	}
	CommitteeConfirmations = storjFlag{
		index:10,
		parameter:0,
		dimension:2,
	}
	OneTimeAddr = storjFlag{
		index:11,
		parameter:0,
	}
	OneTimeAddrConfirmed = storjFlag{
		index:12,
		parameter:0,
	}
	CertificateAddr = storjFlag{
		index:13,
		parameter:0,
	}
	confirmedMainAddress = storjFlag{
		index:14,
		parameter:0,
	}
	confirmedSubAddress = storjFlag{
		index:15,
		parameter:0,
		returns:"",
	}
)

func main()  {
	//posLen(50)

	//expandTo(confirmedMainAddress,"")
	index := "0000000000000000000000000000000000000000000000000000000000000006"
	key :=  "00000000000000000000000xbccc714d56bc0da0fd33d96d2a87b680dd6d0df6"
	fmt.Println("key+index:", key + index)
	fmt.Println([]byte(key+index))
	a := crypto.Keccak256([]byte(key+index))
	fmt.Println("a hash:", a, "\n string hash: ", fmt.Sprintf("%x\n", a))
	//increaseHexByOne(a)
}

//type Hash [HashLength]byte

func expandTo(methods storjFlag, key string)  {
	//newHash[len(newHash)-1] = methods.index
	//indexed := fmt.Sprintf("%x\n", newHash)
	//fmt.Printf("%s",indexed)
	//newHash1:=make([]byte,len(newHash))
	//copy(newHash1[:],newHash[:])
	//f := crypto.Keccak256(newHash1)
	newHash := make([]byte,common.HashLength)
	newHash[len(newHash)-1] = byte(methods.index)
	indexed := fmt.Sprintf("%x\n", newHash)
	fmt.Println(indexed)
	//f := crypto.Keccak256(newHash)
	//fmt.Printf("%s", fmt.Sprintf("%x\n", f))

	//	1.string caSign;
	//    string certMsg;
	//    string pubKey;

	//	2.



}

func posLen(pos uint64) string {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, pos)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//fmt.Printf("Encoded: % x\n", buf.Bytes())

	posLen := fmt.Sprintf("%x\n", buf.Bytes())
	fmt.Println(":::::", posLen)

	return string(posLen)
}

func increaseHexByOne(indexKeyHash []byte) string {
	//check := make([]uint64, 3)
	//rbuf := bytes.NewReader(indexKeyHash)
	//err := binary.Read(rbuf, binary.LittleEndian, &check)
	//if err != nil {
	//	fmt.Println("binary.Read failed:", err)
	//}
	//fmt.Printf("Decoded: %v\n", check)
	fmt.Println(hex.EncodeToString(indexKeyHash))
	//a := fmt.Sprintf("%b", indexKeyHash)

	return ""
}
