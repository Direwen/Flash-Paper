package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func Encrypt(plainText string) (string, error) {
	// Create GCM (Galois/Counter Mode) wrapper around the cipher
	// GCM provides both encryption AND authentication (prevents tampering)
	gcm, err := getGCM()
	if err != nil {
		return "", err
	}

	// Create a unique nonce (number used once)
	nonce := make([]byte, gcm.NonceSize())
	// Generate random bytes using crypto/rand (secure random source)
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// gcm.Seal prepends the nonce to the encrypted data, then encrypts
	// Format: [nonce][encrypted data][authentication tag]
	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	// Convert binary data to base64 string for easy storage/transmission
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(cryptoText string) (string, error) {

	// Decode base64 string back to binary data
	cipherText, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	// Get the GCM cipher instance
	gcm, err := getGCM()
	if err != nil {
		return "", err
	}

	// Get the nonce size (typically 12 bytes for GCM)
	nonceSize := gcm.NonceSize()

	// Ensure the ciphertext is long enough to contain the nonce
	if len(cipherText) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	// Extract the nonce from the beginning and the encrypted data from the rest
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	// Decrypt the data using GCM
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	// Convert decrypted bytes back to string
	return string(plainText), nil
}

// getGCM creates and returns a Galois/Counter Mode (GCM) cipher for AES encryption
func getGCM() (cipher.AEAD, error) {
	// Load the 32-byte encryption key from environment variable
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(key) != 32 {
		return nil, errors.New("invalid key size: must be 32 bytes")
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Return GCM mode cipher for authenticated encryption
	return cipher.NewGCM(block)
}
