package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	hexString := ""
	xorAgainst := ""

	fmt.Printf("Enter First String ")
	fmt.Scanf("%s\n", &hexString)

	fmt.Printf("String to XOR Against ")
	fmt.Scanf("%s\n", &xorAgainst)

	if len(hexString) != len(xorAgainst) {
		log.Println("String length is not equal")
		time.Sleep(2 * time.Second)
		return
	}

	hexBigInt := &big.Int{}
	xorAgainstBigInt := &big.Int{}

	hexBigInt.SetString(hexString, 16)
	xorAgainstBigInt.SetString(xorAgainst, 16)

	fmt.Printf("%s\n", hexBigInt.Xor(hexBigInt, xorAgainstBigInt).String())
}
