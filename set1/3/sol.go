package main

import (
	"encoding/hex"
	"strings"
	"fmt"
)

func main() {
	results := []int{}
	answer := ""
	hexstr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	//Decode hex
	dec, _ := hex.DecodeString(hexstr)

	//Brute force, XOR it with every byte from 0-255 and then pick the correct answer i.e. the answer with nicest/common english character
	//One way to select best string is chi-square test, or just count the max number of spaces which can work in this case but not the most appropriate way to select correct answer
	for i := 0; i < 255; i++ {
		output := ""
		for _, r := range dec {
			output += string(r ^ byte(i))
		}

		spaces := len(strings.Split(output, " "))

		isMax := maxInArray(results, spaces)

		if isMax {
			answer = output
		}

		results = append(results, spaces)
	}

	fmt.Println(answer)
}

func maxInArray(arr []int, key int) bool {

	for i := 0; i < len(arr); i++ {
		if arr[i] >= key {
			return false
		}
	}

	return true
}
