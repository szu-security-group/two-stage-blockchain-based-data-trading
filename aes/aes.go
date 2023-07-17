package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
)

// 加密函数
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	// 创建一个新的 AES 密码分组，使用提供的密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个新的 GCM 加密器
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 创建一个随机 nonce（加密过程中使用的初始化向量）
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// 加密明文并将 nonce 附加到密文开头
	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

// 解密函数
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	// 创建一个新的 AES 密码分组，使用提供的密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个新的 GCM 加密器
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 从密文中提取 nonce
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 解密密文
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
func ReadKey(keyPath string) ([]byte, error) {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	return key, nil
}
