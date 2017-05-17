package main

import (
	"fmt"
	"strconv"
)

func main() {
	firstString := ""
	// secondString := ""

	fmt.Println("Enter First String")
	fmt.Scanf("%s", &firstString)

	// fmt.Println("Enter Second String")
	// fmt.Scanf("%s", &secondString)

	// if len(firstString) != len(secondString) {
	// 	log.Fatalln("String lengths are not equal!")
	// }

	fmt.Println(getStringInBinary(firstString))
	// getStringInBinary(secondString)
}

func getStringInBinary(str string) []string {
	bitsArray := make([]string, len(str))
	for index, char := range []byte(str) {
		charInBits := strconv.FormatInt(int64(char), 2)
		bitsArray[index] = charInBits
	}

	return bitsArray
}
