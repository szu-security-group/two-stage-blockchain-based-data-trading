package hmac

import (
	"crypto/sha256"
)

const (
	opad      = 0x5c
	ipad      = 0x36
	blockSize = 32
)

func HMACK(key []byte, message []byte) []byte {
	// XOR key with opad and ipad
	keyPadded := make([]byte, blockSize)
	copy(keyPadded, key)
	for i := range keyPadded {
		keyPadded[i] ^= opad
	}
	keyIpad := make([]byte, blockSize)
	copy(keyIpad, key)
	for i := range keyIpad {
		keyIpad[i] ^= ipad
	}

	// Hash key XORed with ipad concatenated with the message
	h := sha256.New()
	h.Write(keyIpad)
	h.Write(message)
	innerHash := h.Sum(nil)

	// Hash key XORed with opad concatenated with the inner hash
	h = sha256.New()
	h.Write(keyPadded)
	h.Write(innerHash)
	finalHash := h.Sum(nil)

	return finalHash
}
