package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"crypto/ecdsa"
	"encoding/hex"
)

func main()  {

	R := make([]ecdsa.PublicKey, 5)
	for i := 0; i < 5; i++ {
		key, err := crypto.GenerateKey()
		if err !=nil {
			fmt.Println(err)
		}
		privateKey := hex.EncodeToString(key.D.Bytes())
		fmt.Println("private----->",privateKey)

		pub:=common.ToHex(crypto.FromECDSAPub(&key.PublicKey))
		fmt.Println("pub------>",pub)


		R[i].X, R[i].Y = crypto.S256().ScalarBaseMult(key.D.Bytes())
		R[i].Curve = crypto.S256()

		fmt.Printf("X: %x\n",R[i].X)
		fmt.Printf("Y: %x\n",R[i].Y)

		//fmt.Println(common.ToHex(crypto.FromECDSAPub(&R[i])))
	}
}

/*
private-----> 2bdc8a116b87c76a476af0dc85a13fdecb41ea3a0f08eb803ed962e4032d4b2f
pub------> 0x048da88793ee80f81662407f1dc82522e1f3d89698722063cd5a2c296eab7c4f4350d6711d681cfbe64d3ef610631b9655215f12916af812259acde908c7e6325e
X: 8da88793ee80f81662407f1dc82522e1f3d89698722063cd5a2c296eab7c4f43
Y: 50d6711d681cfbe64d3ef610631b9655215f12916af812259acde908c7e6325e
private-----> 2c6652d1392779fe6ccc0542ef640710621fa795b7260bdaed679f07c5b299d3
pub------> 0x04284bdbba0e8de36af91ea11d7b313ca962ab560a8fb15e01f92df73c4be8102051c566d7b2595cb1cefbb199e9c68aad9ba5ef885e19da390f19b1d66543638d
X: 284bdbba0e8de36af91ea11d7b313ca962ab560a8fb15e01f92df73c4be81020
Y: 51c566d7b2595cb1cefbb199e9c68aad9ba5ef885e19da390f19b1d66543638d
private-----> 916d985d7f5d49e9519258c387854a868570eb8f4661b744f68ce96ade1d459f
pub------> 0x04ec2ed1e8f962b3d1eb3455bef3074dcfc0951f2fbc62eeb4275fe9451bd960d10d1a73935a914ac4f5ff69d972d5b5a4382c6f1d44e90aa9bc36fabdbe99e4dc
X: ec2ed1e8f962b3d1eb3455bef3074dcfc0951f2fbc62eeb4275fe9451bd960d1
Y: d1a73935a914ac4f5ff69d972d5b5a4382c6f1d44e90aa9bc36fabdbe99e4dc
private-----> 8acbdf713f9f79b2743acd1996a3ae5a8e8caf7a7a1417e30320bda4cddaa879
pub------> 0x044a04465180a94942f89b116d6f05daa392392eaa38dd007f658ff25d5cfe045c1f8a294cef76e44dba8f0e10ecf9e6ef05f1531f895585e078876844dfef432a
X: 4a04465180a94942f89b116d6f05daa392392eaa38dd007f658ff25d5cfe045c
Y: 1f8a294cef76e44dba8f0e10ecf9e6ef05f1531f895585e078876844dfef432a
private-----> a4a1cda97c1ac646cf0134e3ab33653826b5d8baedb18e6fbf8d16d5385ca9eb
pub------> 0x042e8d0966740a2e73be5bd5cc862e2898c716cf2fb8e23403ea602ee2b53044fc550a46dbf4baf1f43d7050414c89dfd3210086396d00e89c78ed63cc8ea95ed4
X: 2e8d0966740a2e73be5bd5cc862e2898c716cf2fb8e23403ea602ee2b53044fc
Y: 550a46dbf4baf1f43d7050414c89dfd3210086396d00e89c78ed63cc8ea95ed4
*/