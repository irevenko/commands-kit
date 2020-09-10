package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

func decrypt(encryptedString string, keyString string) (decryptedString string) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plainText)
}

func main() {
	bytes := []byte{222, 158, 110, 144, 26, 231, 109, 23, 128, 200, 132, 158, 84, 38, 149, 145, 59, 9, 50, 213, 12, 47, 164, 22, 199, 65, 131, 86, 189, 54, 144, 198}
	key := hex.EncodeToString(bytes)

	fmt.Println("Enter your encrypted message:")
	input := bufio.NewReader(os.Stdin)
	text, _ := input.ReadString('\n')

	decrypted := decrypt(text, key)
	fmt.Printf("decrypted message: %s\n", decrypted)
}
