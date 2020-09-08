package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your message:")
	input := bufio.NewReader(os.Stdin)
	text, _ := input.ReadString('\n')

	hash := sha256.New()
	hash.Write([]byte(text))

	sha_hash := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("Encoded data:\n", sha_hash)
}
