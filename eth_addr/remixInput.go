//@Time  : 2018/7/26 23:44
//@Author: Greg Li
package main

import (
	"fmt"
	"encoding/hex"
	"time"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"go-wanchain-feature-scan/ethclient"
	"reflect"
)

func main()  {

}
func raw(client *ethclient.Client){
	address := contractAddress
	d := time.Now().Add(1000 * time.Millisecond)
	tx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	testabi, err := abi.JSON(strings.NewReader(TestABI))
	unlockedKey, err := keystore.DecryptKey([]byte(key), "password")
	nonce, _ := client.NonceAt(ctx, unlockedKey.Address, nil)

	if err !=nil {
		fmt.Println("wrong passcode")
	} else{
		fmt.Println(nouce)
		if err !=nil {
			fmt.Println(err)
		} else{
			bytesData, _ := testabi.Pack("<<function name>>", "<<argument1>>", "<<>argument2>")
			tx := types.NewTransaction(nonce, common.HexToAddress(address), nil, big.NewInt(10000000), big.NewInt(0), bytesData)
			signTx, _ := types.SignTx(tx, types.HomesteadSigner{}, unlockedKey.PrivateKey)
			err := client.SendTransaction(ctx, signTx)
			if err !=nil {
				fmt.Println(err)
			} else {
				fmt.Println(signTx)
			}
		}
	}
}

