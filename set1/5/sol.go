package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatingXORCipher())
}

func RepeatingXORCipher() string {

	str := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"

	output := ""
	for i := 3; i < len(str); i += 3 {

		output += string(key[0] ^ str[i-3])
		output += string(key[1] ^ str[i-2])
		output += string(key[2] ^ str[i-1])

	}

	return fmt.Sprintf("%x", output)

}
