package main

import (
	"encoding/base64"
	"fmt"
	"github.com/getlantern/errors"
	"io/ioutil"
	"log"
)

/*
There's a file here. It's been base64'd after being encrypted with repeating-key XOR.

Decrypt it.

Here's how:

Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
Write a function to compute the edit distance/Hamming distance between two strings. The Hamming distance is just the number of differing bits. The distance between:
this is a test
and
wokka wokka!!!
is 37. Make sure your code agrees before you proceed.
For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
The KEYSIZE with the smallest normalized edit distance is probably the key. You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4 KEYSIZE blocks instead of 2 and average the distances.
Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
Now transpose the blocks: make a block that is the first byte of every block, and a block that is the second byte of every block, and so on.
Solve each block as if it was single-character XOR. You already have code to do this.
For each block, the single-byte XOR key that produces the best looking histogram is the repeating-key XOR key byte for that block. Put them together and you have the key.
This code is going to turn out to be surprisingly useful later on. Breaking repeating-key XOR ("Vigenere") statistically is obviously an academic exercise, a "Crypto 101" thing. But more people "know how" to break it than can actually break it, and a similar technique breaks something much more important.
 */

func main() {

	cipherBytes, err := ioutil.ReadFile("cipher.dat")
	if err != nil {
		log.Fatalln("Error occurred in cipher.dat", err.Error())
	}

	decStr, _ := base64.StdEncoding.DecodeString(string(cipherBytes))
	distances := []int{}
	keySize := 0

	//Find the appropriate key
	for i := 2; i < 41; i++ {
		distance, err := normalizedDistance(i, decStr[:i], decStr[i:(2 * i)])
		if err != nil {
			log.Fatalln("Error in calculating distance", err.Error())
		}
		isMin := isMinInArray(distances, distance)
		if isMin {
			keySize = i
		}
		distances = append(distances, distance)
	}
	// Now, We know the keysize at this point

	decStrBlocks := [][]byte{}

	//Divide decoded base64 bytes into blocks of length keysize
	for i := 0; i < len(decStr); i += keySize {
		end := i + keySize
		if end > len(decStr) {
			end = len(decStr)
		}

		decStrBlocks = append(decStrBlocks, decStr[i:end])
	}

	fmt.Println(decStrBlocks)

	//Transpose
	transposed := transpose(decStrBlocks, keySize)
	for i := 0; i < len(transposed); i++ {
		fmt.Println(string(transposed[i]))
	}

	//TODO:Add chi-square test to find out key of each item in transposed array

}

func transpose(decBlocks [][]byte, keySize int) [][]byte {
	newDecBlocks := [][]byte{}

	for i := 0; i < keySize; i++ {
		tempBlock := []byte{}
		for j := 0; j < len(decBlocks); j++ {

			for k := 0; j < len(decBlocks[j]); k++ {
				tempBlock = append(tempBlock, decBlocks[j][i])
			}

		}
		newDecBlocks = append(newDecBlocks, tempBlock)

	}
	return newDecBlocks
}

func normalizedDistance(keySize int, str1, str2 []byte) (int, error) {
	distance, err := calcHammingDistance(str1, str2)
	if err != nil {
		return 0, err
	}
	distance = distance / keySize
	return distance, nil
}

func calcHammingDistance(str1, str2 []byte) (int, error) {
	distance := 0

	if len(str1) != len(str2) {
		return 0, errors.New("Strings of Unequal length")
	}

	bin1 := ""
	bin2 := ""

	for _, b := range str1 {
		bin1 += fmt.Sprintf("%08b", b)
	}

	for _, b := range str2 {
		bin2 += fmt.Sprintf("%08b", b)
	}

	for i := 0; i < len(bin1); i++ {
		res := int(bin1[i]) ^ int(bin2[i])
		if res != 0 {
			distance++
		}
	}

	return distance, nil
}

func isMinInArray(arr []int, key int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] <= key {
			return false
		}
	}
	return true
}
