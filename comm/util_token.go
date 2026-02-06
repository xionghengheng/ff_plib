package comm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// 固定密钥（16字节，用于AES-128加密）
const aesSecretKey = "FF_TRIAL_KEY_16!"

// 生成H5链接token（包含记录ID，使用AES加密+Base64编码）
func GenerateH5LinkToken(recordId int64, createTs int64) string {
	// 使用记录ID+教练ID+时间戳生成原始数据
	data := fmt.Sprintf("%d_%d", recordId, createTs)

	// AES加密
	encrypted, err := aesEncrypt([]byte(data), []byte(aesSecretKey))
	if err != nil {
		Printf("generateH5LinkToken aesEncrypt err:%v\n", err)
		return ""
	}

	// URL安全的Base64编码
	return base64.URLEncoding.EncodeToString(encrypted)
}

// 解密H5链接token，返回recordId, coachId, createTs
func DecryptH5LinkToken(token string) (recordId int64, createTs int64, err error) {
	// Base64解码
	cipherText, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return 0, 0, fmt.Errorf("Base64解码失败: %v", err)
	}

	// AES解密
	plainText, err := aesDecrypt(cipherText, []byte(aesSecretKey))
	if err != nil {
		return 0, 0, fmt.Errorf("AES解密失败: %v", err)
	}

	// 解析原始数据: recordId_createTs
	var rId, ts int64
	_, err = fmt.Sscanf(string(plainText), "%d_%d", &rId, &ts)
	if err != nil {
		return 0, 0, fmt.Errorf("解析数据失败: %v", err)
	}

	return rId, ts, nil
}

// AES加密（CBC模式）
func aesEncrypt(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// PKCS7填充
	blockSize := block.BlockSize()
	padding := blockSize - len(plainText)%blockSize
	padText := make([]byte, len(plainText)+padding)
	copy(padText, plainText)
	for i := len(plainText); i < len(padText); i++ {
		padText[i] = byte(padding)
	}

	// 使用密钥前16字节作为IV
	iv := key[:blockSize]
	encrypted := make([]byte, len(padText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, padText)

	return encrypted, nil
}

// AES解密（CBC模式）
func aesDecrypt(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(cipherText) < blockSize || len(cipherText)%blockSize != 0 {
		return nil, fmt.Errorf("密文长度无效")
	}

	// 使用密钥前16字节作为IV
	iv := key[:blockSize]
	decrypted := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, cipherText)

	// 去除PKCS7填充
	padding := int(decrypted[len(decrypted)-1])
	if padding > blockSize || padding == 0 {
		return nil, fmt.Errorf("填充无效")
	}
	// 校验填充
	for i := len(decrypted) - padding; i < len(decrypted); i++ {
		if decrypted[i] != byte(padding) {
			return nil, fmt.Errorf("填充校验失败")
		}
	}

	return decrypted[:len(decrypted)-padding], nil
}
