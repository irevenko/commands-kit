package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	key, _ := hex.DecodeString(keyString)
	plainText := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)
	return fmt.Sprintf("%x", cipherText)
}

func main() {
	var cmd = &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt any string",
		Long:  "A simple CLI tool for encrypting your data \nSource code: https://github.com/irevenko/commands-kit",
		Run: func(cmd *cobra.Command, args []string) {
			bytes := []byte{222, 158, 110, 144, 26, 231, 109, 23, 128, 200, 132, 158, 84, 38, 149, 145, 59, 9, 50, 213, 12, 47, 164, 22, 199, 65, 131, 86, 189, 54, 144, 198}
			key := hex.EncodeToString(bytes)

			fmt.Println("Enter your message:")
			input := bufio.NewReader(os.Stdin)
			text, _ := input.ReadString('\n')

			encrypted := encrypt(text, key)
			fmt.Printf("Encrypted message: %s\n", encrypted)
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
