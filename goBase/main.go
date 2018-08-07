package main

import (
	"fmt"
	//"io/ioutil"

	//"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	//"github.com/ethereum/go-ethereum/console"
	//"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/log"
	//"gopkg.in/urfave/cli.v1"
)

func AccountGenerate(password string, ks *keystore.KeyStore) error {
	fmt.Printf("Generate new account...")

	account, err := ks.NewAccount(password)
	if err != nil {
		utils.Fatalf("Failed to create account: %v", err)
	}
	fmt.Printf("address : {%x}\n",account.Address)
	return nil
}

func UnluckMinerAccount(passwd string, ks *keystore.KeyStore) (*keystore.Key, error) {
	fmt.Printf("unlocking Account with null passwd.")
	accs := ks.Accounts()
	fmt.Printf("address : {%x}\n",accs[0])
	err := ks.TimedUnlock(accs[0], passwd, 50000000)
	if err != nil {
		utils.Fatalf("Failed to unlock account: %v", err)
		return nil, err
	}
	_, key, err := ks.GetDecryptedKey(accs[0], passwd)
	//fmt.Printf("\nprivatekey : {%x}\n", key.PrivateKey)
	fmt.Printf("unlock success.")
	return key, nil
}

func main()  {
	// 密码设置为空，临时
	passwd := ""
	// new keystore
	ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)
	// 产生账户文件，目录为" ./ "
	AccountGenerate(passwd, ks)
	// 根据密码，读取账户文件
	key, err := UnluckMinerAccount("", ks)
	if err != nil {
		utils.Fatalf("Failed to unlock account: %v", err)
	}
	// 根据私钥，签名交易
	fmt.Printf("\nprivatekey : {%x}\n", key.PrivateKey)
}
