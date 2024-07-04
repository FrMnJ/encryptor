package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"math/big"
)

func Encrypt(key []byte, file []byte) (encryptedFile []byte, err error) {
	cypherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cypherBlock)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	encryptedFile = gcm.Seal(nonce, nonce, file, nil)
	return encryptedFile, nil
}

func Decrypt(key []byte, file []byte) (decryptedFile []byte, err error) {
	cypherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher block: %w", err)
	}
	gcm, err := cipher.NewGCM(cypherBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}
	nonceSize := gcm.NonceSize()
	if len(file) < nonceSize {
		return nil, fmt.Errorf("encrypted file size is too small")
	}
	nonce, encFile := file[:nonceSize], file[nonceSize:]
	decryptedFile, err = gcm.Open(nil, nonce, encFile, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}
	return decryptedFile, nil
}

func Generate32Key() []byte {
	const lettersNumbersAndSymbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+1234567890"
	key := make([]byte, 32)
	for i := 0; i < 32; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(lettersNumbersAndSymbols))))
		if err != nil {
			panic(err)
		}
		key[i] = byte(lettersNumbersAndSymbols[n.Int64()])
	}
	return key
}
