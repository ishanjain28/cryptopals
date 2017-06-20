package main

import (
	"fmt"
	"bytes"
	"strconv"
)

func main() {

	//message1 := []byte("I go crazy when I hear a cymbal")
	//message1 := []byte("Burning 'em, if you ain't quick and nimble")

	message1 := []byte("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")

	var fullText []byte
	for i := 0; i < len(message1)-2; i += 3 {

		fullText = append(fullText, message1[i]^byte('I'))
		fullText = append(fullText, message1[i+1]^byte('C'))
		fullText = append(fullText, message1[i+2]^byte('E'))
	}
	//answer := "b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	for i := 0; i < len(fullText); i++ {
		//fmt.Printf("%x", fullText[i])

		fmt.Println([]rune(string(fullText[i])))
	}
	fmt.Println(bytes.Runes(fullText))

	fmt.Println(strconv.Itoa(int(fullText[0])))
}
