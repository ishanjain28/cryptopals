package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

func main() {

	inputString := ""

	fmt.Println("Enter Hex String")

	fmt.Scanf("%s\n", &inputString)
	hexString, err := hex.DecodeString(inputString)
	if err != nil {
		log.Printf("Invalid Hex\n")
	}
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(hexString)))

	//Wait a little bit before quitting
	time.Sleep(2 * time.Second)
}
