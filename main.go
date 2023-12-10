package main

import (
	"fmt"

	"github.com/Luftalian/RSA_practice/crypto"
	"github.com/Luftalian/RSA_practice/rsa"
)

func main() {
	var crypto crypto.Crypto = rsa.NewRSA(8429, 9901)

	a := 2023
	fmt.Printf("平文: %d\n", a)

	b, err := crypto.Encrypt(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("暗号文: %d\n", b)

	aPrime, err := crypto.Decrypt(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("復号文: %d\n", aPrime)
}
