package main

import (
	"encoding/base64"
	"fmt"
	"github.com/getlantern/errors"
	"io/ioutil"
	"log"
)

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

	decStrBlocks := [][]byte{}

	//Divide decoded base64 bytes into blocks of length keysize
	for i := 0; i < len(decStr); i += keySize {
		end := i + keySize
		if end > len(decStr) {
			end = len(decStr)
		}

		decStrBlocks = append(decStrBlocks, decStr[i:end])
	}

	//Transpose
	transposed := transpose(decStrBlocks, keySize)
	for i := 0; i < len(transposed); i++ {
		fmt.Println(string(transposed[i]))
	}

	//	TODO:Add chi-square test to find out key of each item in transposed array

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
