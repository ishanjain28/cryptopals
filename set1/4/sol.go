package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"bufio"
	"io"
	"log"
)

/*
Detect single-character XOR
One of the 60-character strings in this file has been encrypted by single-character XOR.

Find it.

(Your code from #3 should help.)
 */

func main() {

	file, err := os.Open("hash.dat")
	if err != nil {
		log.Fatalln("Error in opening hash.dat", err.Error())
	}

	defer file.Close()

	answer := ""
	globalMaxScore := 0
	hashes := bufio.NewReader(file)

	for true {
		line, _, err := hashes.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Fatalln("Error occurred in reading file", err.Error())
			}
			break
		}
		decHash, _ := hex.DecodeString(string(line))

		lanswer := ""
		maxScore := 0

		for j := 0x00; j <= 0xff; j++ {
			output := ""
			score := 0
			for _, r := range decHash {
				xor := r ^ byte(j)

				output += string(xor)
				scoreResult(xor, &score)
			}

			if score > maxScore {
				lanswer = output
				maxScore = score

			}
		}

		if maxScore > globalMaxScore {
			globalMaxScore = maxScore
			answer = lanswer
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
