package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateSHA256Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func main() {
	data := "Hello, World!"
	hash := generateSHA256Hash(data)

	fmt.Println("Original Data:", data)
	fmt.Println("SHA256 Hash:", hash)
}
