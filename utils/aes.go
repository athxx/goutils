package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

const defaultKey = `4)~Z-GyoambQs#!}`

func AesEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	rawData = pKCS5Padding(rawData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(rawData))
	blockMode.CryptBlocks(encrypted, rawData)
	return encrypted, nil
}

func AesDecrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	rawData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(rawData, encrypted)
	rawData = pKCS5UnPadding(rawData)
	return rawData, nil
}

// AesCrypt this struct can customize 'IV' instead of using key size be the IV
// NOTICE: iv length must equal key block size
type AesCrypt struct {
	Key []byte
	Iv  []byte
}

func (a *AesCrypt) Encrypt(data []byte) ([]byte, error) {
	if len(a.Key) != len(a.Iv) {
		return nil, errors.New(`iv length must equal key length`)
	}
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	content := pKCS5Padding(data, block.BlockSize())
	cipherBytes := make([]byte, len(content))
	encoder := cipher.NewCBCEncrypter(block, a.Iv)
	encoder.CryptBlocks(cipherBytes, content)
	return cipherBytes, nil
}

func (a *AesCrypt) Decrypt(src []byte) (data []byte, err error) {
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	decoder := cipher.NewCBCDecrypter(block, a.Iv)
	decoder.CryptBlocks(decrypted, src)
	return pKCS5UnPadding(decrypted), nil
}

func pKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	return append(cipherText, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pKCS5UnPadding(encrypt []byte) []byte {
	return encrypt[:len(encrypt)-int(encrypt[len(encrypt)-1])]
}
