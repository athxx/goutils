package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// 生成私钥文件 TODO 未指定路径
func RsaKeyGen(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  `RSA PRIVATE KEY`,
		Bytes: derStream,
	}
	priFile, err := os.Create(`private.pem`)
	if err != nil {
		return err
	}
	err = pem.Encode(priFile, block)
	priFile.Close()
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  `PUBLIC KEY`,
		Bytes: derPkix,
	}
	pubFile, err := os.Create(`public.pem`)
	if err != nil {
		return err
	}
	err = pem.Encode(pubFile, block)
	pubFile.Close()
	if err != nil {
		return err
	}
	return nil
}

// 生成私钥文件, 返回 privateKey , publicKey, error
func RsaKeyGenText(bits int) (string, string, error) { // bits 字节位 1024/2048
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return ``, "", err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  `RSA PRIVATE KEY`,
		Bytes: derStream,
	}
	priBuff := bytes.NewBuffer(nil)
	err = pem.Encode(priBuff, block)
	if err != nil {
		return ``, ``, err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return ``, ``, err
	}
	block = &pem.Block{
		Type:  `PUBLIC KEY`,
		Bytes: derPkix,
	}
	pubBuff := bytes.NewBuffer(nil)
	err = pem.Encode(pubBuff, block)
	if err != nil {
		return ``, ``, err
	}
	return priBuff.String(), pubBuff.String(), nil
}

// 加密
func RsaEncrypt(rawData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New(`public key error`)
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, rawData)
}

// 解密
func RsaDecrypt(cipherText, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New(`private key error`)
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}
