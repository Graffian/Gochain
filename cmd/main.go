package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"strconv"
)

type Block struct{
	Index int
	Timestamp int
	Data string
	PrevHash string
	Nonce int
	Hash string
}

const balance int = 10000

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
	fmt.Printf("---------------Your balance is %d GC------------------ \n" , balance)
	fmt.Print("Make a Transaction (y/n): ")
	conf,err := reader.ReadString('\n')
	if err!=nil{
		fmt.Println("An err occured: " , err)
		return
	}
	conf = strings.TrimSpace(conf)



	if conf == "y"{
		fmt.Printf("Name of the person you want to send it to : ")
		recvr,_ := reader.ReadString('\n')
		recvr = strings.TrimSpace(recvr)
		fmt.Printf("Amount you want to send: ")
		amt,_ := reader.ReadString('\n')
		recvr_amt,_ := strconv.Atoi(strings.TrimSpace(amt))
		if (recvr_amt > balance && recvr_amt<=0){
			fmt.Println("Not enough balance")
		}else{
			fmt.Printf("You sent %d GC to %s" , recvr_amt , recvr)
		}
	}




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
