package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type Block struct{
	Index int
	Timestamp int
	Data string
	PrevHash string
	Nonce int
	Hash string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name :  ")
	text,err := reader.ReadString('\n')
	if err!=nil{
		fmt.Println("An error occured: " , err)
		return
	}
	text = strings.TrimSpace(text)
	fmt.Println("<<<<<<<<<<<<<<<<< WELCOME TO GOCHAIN >>>>>>>>>>>>>>>>>>>>>>>>>")
	keys := gen_keys()
	fmt.Println("Your public key is : " , keys["publicKey"])
	fmt.Println("Your secret key is : " , keys["privateKey"])
}

func gen_keys() map[string]any{
	privateKey,err := ecdsa.GenerateKey(elliptic.P256() , rand.Reader)
	if err != nil{
		fmt.Println("Error in generating key: " , err)
		return nil
	}
	publicKey := privateKey.PublicKey
	keys := map[string]any{
		"publicKey" : publicKey.X,
		"privateKey" : privateKey.D,
	}
	return keys

}
