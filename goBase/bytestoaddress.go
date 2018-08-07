package main

import (
	"math/big"
	"fmt"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/sha3"

)

const (
	HashLength     = 32
	AddressLength  = 20
	WAddressLength = 66
)

// Address represents the 20 byte address of an ordinary account.
type Address [AddressLength]byte

// Hash represents the 32 byte Keccak256 hash of arbitrary data.
type Hash [HashLength]byte

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

// Sets the address to the value of b. If b is larger than len(a) it will panic
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

func main()  {
	var otaBalanceStorageAddr = BytesToAddress(big.NewInt(300).Bytes())
	fmt.Println(otaBalanceStorageAddr.Hex())

	var otaImageStorageAddr   = BytesToAddress(big.NewInt(301).Bytes())
	fmt.Println(otaImageStorageAddr)
}

// Hex returns an EIP55-compliant hex string representation of the address.
func (a Address) Hex() string {
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