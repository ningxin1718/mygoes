package main

import (
	"log"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"encoding/hex"
)


func main() {
	//path, _ := filepath.Abs("./resources/etherdelta.abi")
	//file, err := ioutil.ReadFile(path)
	//if err != nil {
	//	fmt.Println("Failed to read file:", err)
	//}
	//
	//
	//edabi, err := abi.JSON(strings.NewReader(string(file)))
	//if err != nil {
	//	fmt.Println("Invalid abi:", err)
	//}
	edabi, err := abi.JSON(strings.NewReader("[{\"constant\":true,\"inputs\":[],\"name\":\"unConfirmedAddressLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCommittee\",\"outputs\":[{\"name\":\"added\",\"type\":\"bool\"},{\"name\":\"execution\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"CommitteePublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"confirmedSubAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"confirmedMainAddressLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"confirmedMainAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_COMMITTEEMAN_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkOneTimeAddrAdded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requirement\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkAddrConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"CommitteeConfirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unConfirmedAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"CertificateAddr\",\"outputs\":[{\"name\":\"added\",\"type\":\"bool\"},{\"name\":\"confirmed\",\"type\":\"bool\"},{\"name\":\"addressType\",\"type\":\"uint8\"},{\"name\":\"ringSig\",\"type\":\"string\"},{\"name\":\"pubSKey\",\"type\":\"string\"},{\"name\":\"publicKeyMirror\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"OneTimeAddr\",\"outputs\":[{\"name\":\"confirmed\",\"type\":\"bool\"},{\"name\":\"caSign\",\"type\":\"string\"},{\"name\":\"certMsg\",\"type\":\"string\"},{\"name\":\"pubKey\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CertToAddress\",\"outputs\":[{\"name\":\"confirmed\",\"type\":\"bool\"},{\"name\":\"toAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"OneTimeAddrConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CMMTTEEs\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"OneTimeAddrConfirmedLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"confirmedSubAddressLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"certIDCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"Committeeman\",\"type\":\"address\"}],\"name\":\"CommitteemanAddition\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"_createrPubKey\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newPending\",\"type\":\"address\"}],\"name\":\"removeCommittee\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"submitIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"added\",\"type\":\"bool\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_certID\",\"type\":\"uint256\"},{\"name\":\"_confirm\",\"type\":\"bool\"}],\"name\":\"confirmCert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addressType\",\"type\":\"uint8\"},{\"name\":\"_ringSig\",\"type\":\"string\"},{\"name\":\"_pub_S_Key\",\"type\":\"string\"},{\"name\":\"_publicKeyMirror\",\"type\":\"string\"}],\"name\":\"summitCert\",\"outputs\":[{\"name\":\"_certID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ringSig\",\"type\":\"string\"},{\"name\":\"_pub_S_Key\",\"type\":\"string\"},{\"name\":\"_publicKeyMirror\",\"type\":\"string\"}],\"name\":\"storeSubUserCert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pubkey\",\"type\":\"string\"},{\"name\":\"_sign\",\"type\":\"string\"},{\"name\":\"_CA\",\"type\":\"string\"}],\"name\":\"storeOneTimeAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ringSig\",\"type\":\"string\"},{\"name\":\"_pub_S_Key\",\"type\":\"string\"},{\"name\":\"_publicKeyMirror\",\"type\":\"string\"}],\"name\":\"storeMainUserCert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"submitIndex\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newPending\",\"type\":\"address\"},{\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"addCommittee\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"Committeeman\",\"type\":\"address\"}],\"name\":\"CommitteemanRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"confirmed\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"submitIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"added\",\"type\":\"bool\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"submitIndex\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"}]"))
	if err != nil {
		log.Fatal(err)
	}

	var orderStruct struct {
		TokenGet   common.Address
		AmountGet  *big.Int
		TokenGive  common.Address
		AmountGive *big.Int
		Expires    *big.Int
		Nonce      *big.Int
		User       common.Address
	}

	bytesData, _ := edabi.Pack("storeMainUserCert", "abc","zxcv","asdf")
	fmt.Println("hex.EncodeToString(a): ",hex.EncodeToString(bytesData))
	fmt.Println(bytesData)


	err = edabi.Unpack(&orderStruct, "storeMainUserCert", bytesData)
	if err != nil {
		fmt.Println("Failed to unpack:", err)
	}


	fmt.Println("TokenGet:", orderStruct.TokenGet.Hex())
}