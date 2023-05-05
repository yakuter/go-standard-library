package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	// Generate a new RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Create a message to sign
	message := []byte("Hello, world!")

	// Sign the message using SHA256 and the private key
	hash := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	}

	// Verify the signature using the public key
	publicKey := &privateKey.PublicKey
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return
	}

	fmt.Println("Signature verification succeeded!")
	fmt.Println("Message:", string(message))
	fmt.Println("Signature:", base64.StdEncoding.EncodeToString(signature))
}
