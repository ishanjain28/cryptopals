package main

import
(
	"fmt"
	"log"
	"encoding/base64"
	"encoding/hex"
)

func main() {
	
	choice := 1
	inputString := ""

	fmt.Println("Choose One Option")
	fmt.Println("1. hex to base64")
	fmt.Println("2. base64 to hex")
	fmt.Println("3. hex to string")
	fmt.Println("4. base64 to string")

	fmt.Printf("Choice: ")
	fmt.Scanf("%d\n", &choice)
	
	switch choice {
	case 1: 
		fmt.Printf("Enter String: ")		
		fmt.Scanf("%s", &inputString)
		fmt.Println(convertToBase64(inputString))
	case 2: 
		fmt.Printf("Enter String: ")
		fmt.Scanf("%s", &inputString)
		fmt.Println(convertToHex(inputString))
	case 3:
		fmt.Printf("Enter Hex: ")
		fmt.Scanf("%s", &inputString)
		fmt.Println(hexToString(inputString))
	case 4:
		fmt.Printf("Enter base64 string: ")
		fmt.Scanf("%s", &inputString)
		fmt.Println(base64ToString(inputString))
	default:
		fmt.Println("Wrong Choice, Please Enter a Valid Choice")
	}
}

func convertToBase64(hexstr string) string {
	hexString, err := hex.DecodeString(hexstr)
	if err != nil {
		log.Fatal("Invalid Hex")
	}
	return base64.StdEncoding.EncodeToString([]byte(hexString))
}

func convertToHex(base64 string) string {
	return hex.EncodeToString([]byte(base64))
}

func hexToString(hexString string) strin4g {
	hexStr, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal("Invalid hex")
	}
	return string(hexStr)
}

func base64ToString(base64Str string) string {
	base64str, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		log.Fatal("Invalid base64 string")
	}
	return string(base64str)
}
