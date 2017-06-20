package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	results := []string{}
	hexstr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	dec, _ := hex.DecodeString(hexstr)

	for i := 0; i < 255; i++ {
		output := ""
		for _, r := range dec {
			output += string(r ^ byte(i))
		}
		//if i == 88 {
		//	fmt.Printf("%x %b %s\n", i, byte(i), output)
		//}

		results = append(results, output)
	}

	fmt.Println(results[120])
}
