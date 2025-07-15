package tcgaming

import (
	"bytes"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

// DESEncrypt handles DES encryption/decryption
type DESEncrypt struct {
	key []byte
}

// NewDESEncrypt creates a new DES encryptor
func NewDESEncrypt(key string) *DESEncrypt {
	// Ensure key is exactly 8 bytes for DES
	keyBytes := []byte(key)
	if len(keyBytes) < 8 {
		// Pad with zeros
		padded := make([]byte, 8)
		copy(padded, keyBytes)
		keyBytes = padded
	} else if len(keyBytes) > 8 {
		// Truncate
		keyBytes = keyBytes[:8]
	}
	
	return &DESEncrypt{
		key: keyBytes,
	}
}

// Encrypt encrypts the input string using DES
func (d *DESEncrypt) Encrypt(input string) (string, error) {
	// For compatibility with Java SDK, we'll use ECB mode
	// Java's default DES without mode specification uses ECB
	plaintext := []byte(input)
	
	// Pad the plaintext
	plaintext = pkcs5Pad(plaintext, des.BlockSize)
	
	// Create cipher
	block, err := des.NewCipher(d.key)
	if err != nil {
		return "", err
	}
	
	// Encrypt using ECB mode
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += des.BlockSize {
		block.Encrypt(ciphertext[i:i+des.BlockSize], plaintext[i:i+des.BlockSize])
	}
	
	// Base64 encode and remove whitespace (matching Java SDK behavior)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	encoded = strings.ReplaceAll(encoded, " ", "")
	encoded = strings.ReplaceAll(encoded, "\n", "")
	encoded = strings.ReplaceAll(encoded, "\r", "")
	encoded = strings.ReplaceAll(encoded, "\t", "")
	
	return encoded, nil
}

// Decrypt decrypts the input string using DES
func (d *DESEncrypt) Decrypt(input string) (string, error) {
	// Base64 decode
	ciphertext, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	
	// Create cipher
	block, err := des.NewCipher(d.key)
	if err != nil {
		return "", err
	}
	
	// Decrypt using ECB mode
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += des.BlockSize {
		block.Decrypt(plaintext[i:i+des.BlockSize], ciphertext[i:i+des.BlockSize])
	}
	
	// Remove padding
	plaintext, err = pkcs5Unpad(plaintext)
	if err != nil {
		return "", err
	}
	
	return string(plaintext), nil
}

// pkcs5Pad pads the data to a multiple of blockSize
func pkcs5Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// pkcs5Unpad removes PKCS5 padding
func pkcs5Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("invalid padding: empty data")
	}
	
	padding := int(data[length-1])
	if padding > length || padding == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	
	// Verify padding bytes
	for i := length - padding; i < length; i++ {
		if data[i] != byte(padding) {
			return nil, fmt.Errorf("invalid padding bytes")
		}
	}
	
	return data[:length-padding], nil
}

// generateIV generates a random IV for CBC mode (if needed in future)
func generateIV() ([]byte, error) {
	iv := make([]byte, des.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil
}