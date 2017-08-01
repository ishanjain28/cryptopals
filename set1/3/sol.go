package main

import (
	"encoding/hex"
	"fmt"
)

/*
Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

Achievement Unlocked
You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.
*/

func main() {
	//results := map[string]int{}
	hexstr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	//Decode hex
	dec, _ := hex.DecodeString(hexstr)
	maxScore := 0
	answer := ""
	//Brute force, XOR it with every byte from 0-255 and then pick the correct answer i.e. the answer with nicest/common english character
	//One way to select best string is chi-square test, or just count the max number of spaces which can work in this case but not the most appropriate way to select correct answer

	for i := 0x00; i < 0x255; i++ {
		score := 0
		output := ""

		for _, r := range dec {
			xor := r ^ byte(i)
			output += string(xor)
			scoreResult(xor, &score)
		}

		if score > maxScore {
			answer = output
			maxScore = score
		}
	}

	fmt.Println(answer)
}

func scoreResult(char byte, score *int) {

	c := []rune(string(char))[0]

	//Check if letters are valid english alphabets, If they are, Increase the score
	if (c >= 97 && c <= 122) || (c <= 90 && c >= 65) || c == 32 {
		*score++
	}
}
