package main

import (
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"os"
)

func main() {
	fmt.Println("please input mnemonic...")
	if len(os.Args) < 1 {
		panic("input error")
	}
	mnemonic := os.Args[1]
	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		panic(err)
	}
	account, err := wallet.Derive(hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0"), false)
	if err != nil {
		panic(err)
	}

	privateKeyHex, _ := wallet.PrivateKeyHex(account)
	fmt.Println("prikey: " + privateKeyHex)
	fmt.Println("address: " + account.Address.Hex())
}
