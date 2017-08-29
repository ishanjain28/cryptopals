package main

import (
	"io/ioutil"
	"log"
	"crypto/aes"
	"fmt"
	"encoding/base64"
)

/*

AES in ECB mode
The Base64-encoded content in this file has been encrypted via AES-128 in ECB mode under the key

"YELLOW SUBMARINE".
(case-sensitive, without the quotes; exactly 16 characters; I like "YELLOW SUBMARINE" because it's exactly 16 bytes long, and now you do too).

Decrypt it. You know the key, after all.

Easiest way: use OpenSSL::Cipher and give it AES-128-ECB as the cipher.

 */

func main() {
	enc, err := ioutil.ReadFile("cipher.dat")
	if err != nil {
		log.Fatalf("error in opening file: %v", err)
	}

	bufenc, err := base64.StdEncoding.DecodeString(string(enc))
	if err != nil {
		log.Fatalf("error in decoding base64 string: %v", err)
	}

	acipher, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		log.Fatalf("error in creating aes cipher: %v", err)
	}

	for i := 0; i < len(bufenc); i += 16 {
		buf := make([]byte, 16)
		acipher.Decrypt(buf, bufenc[i:i+16])
		fmt.Printf("%s", string(buf))
	}

}

