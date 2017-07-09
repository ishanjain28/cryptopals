package main

import (
	"fmt"
	"log"
	"time"
	"encoding/hex"
	"strings"
	"strconv"
)

/*
Fixed XOR
Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965
... should produce:

746865206b696420646f6e277420706c6179
 */

func main() {
	hexString := "1c0111001f010100061a024b53535009181c"
	xorAgainst := "686974207468652062756c6c277320657965"

	//Check if hashes are of equal length, Quit if they are not
	if len(hexString) != len(xorAgainst) {
		log.Println("String length is not equal")
		time.Sleep(2 * time.Second)
		return
	}

	//Decode hexes
	hex1, _ := hex.DecodeString(hexString)
	hex2, _ := hex.DecodeString(xorAgainst)
	bin1 := ""
	bin2 := ""

	//Convert everything to binary.
	//Never work with strings directly, Work with bytes or
	for _, b := range hex1 {
		bin1 += fmt.Sprintf("%08b ", b)
	}

	for _, b := range hex2 {
		bin2 += fmt.Sprintf("%08b ", b)
	}
	bin1Split := strings.Split(bin1, " ")
	bin2Split := strings.Split(bin2, " ")

	//len(bin1Split)-1 here, because one extra empty item is added to slice
	for i := 0; i < len(bin1Split)-1; i++ {
		xor := ""
		for j := 0; j < 8; j++ {
			xor += strconv.Itoa(int(bin1Split[i][j] ^ bin2Split[i][j]))
		}
		val, _ := strconv.ParseInt(xor, 2, 64)
		fmt.Printf("%x", string(val))
	}
	fmt.Println()
}
