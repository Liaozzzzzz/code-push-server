package crypto

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/liaozzzzzz/code-push-server/internal/config"
)

// AESCrypto AES 加解密工具结构体
type AESCrypto struct {
	key []byte
}

// NewAESCrypto 从配置文件创建新的 AES 加解密实例
func NewAESCrypto() (*AESCrypto, error) {
	key := config.C.Security.EncryptionKey
	if len(key) != 32 {
		return nil, errors.New("AES key must be 32 bytes long")
	}
	return &AESCrypto{
		key: []byte(key),
	}, nil
}

// pkcs7Padding PKCS7 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	if padding == 0 {
		padding = blockSize
	}
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	return append(data, padtext...)
}

// pkcs7Unpadding PKCS7 去填充
func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("data is empty")
	}

	unpadding := int(data[length-1])
	if unpadding == 0 || unpadding > length || unpadding > aes.BlockSize {
		return nil, errors.New("invalid padding")
	}

	// 检查填充是否正确
	for i := length - unpadding; i < length; i++ {
		if data[i] != byte(unpadding) {
			return nil, errors.New("invalid padding")
		}
	}

	return data[:(length - unpadding)], nil
}

// encryptBytes 内部加密方法 - ECB 模式
func (a *AESCrypto) encryptBytes(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data cannot be empty")
	}

	// 创建 AES cipher
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// PKCS7 填充
	paddedData := pkcs7Padding(data, aes.BlockSize)

	// ECB 模式加密
	ciphertext := make([]byte, len(paddedData))
	for i := 0; i < len(paddedData); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:i+aes.BlockSize], paddedData[i:i+aes.BlockSize])
	}

	return ciphertext, nil
}

// decryptBytes 内部解密方法 - ECB 模式
func (a *AESCrypto) decryptBytes(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, errors.New("ciphertext cannot be empty")
	}

	// 检查数据长度必须是块大小的倍数
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext length must be multiple of block size")
	}

	// 创建 AES cipher
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// ECB 模式解密
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(plaintext[i:i+aes.BlockSize], ciphertext[i:i+aes.BlockSize])
	}

	// 去除 PKCS7 填充
	unpaddedData, err := pkcs7Unpadding(plaintext)
	if err != nil {
		return nil, fmt.Errorf("failed to remove padding: %w", err)
	}

	return unpaddedData, nil
}

// Encrypt 加密数据
// 使用 AES-256-ECB 模式，返回 base64 编码的字符串
func (a *AESCrypto) Encrypt(plaintext string) (string, error) {
	if plaintext == "" {
		return "", errors.New("plaintext cannot be empty")
	}

	// 加密字节数据
	ciphertext, err := a.encryptBytes([]byte(plaintext))
	if err != nil {
		return "", err
	}

	// 返回 base64 编码的结果
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密数据
// 输入 base64 编码的加密字符串，返回原始明文
func (a *AESCrypto) Decrypt(ciphertext string) (string, error) {
	if ciphertext == "" {
		return "", errors.New("ciphertext cannot be empty")
	}

	// base64 解码
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	// 解密字节数据
	plaintext, err := a.decryptBytes(data)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// EncryptBytes 加密字节数据
func (a *AESCrypto) EncryptBytes(data []byte) ([]byte, error) {
	return a.encryptBytes(data)
}

// DecryptBytes 解密字节数据
func (a *AESCrypto) DecryptBytes(ciphertext []byte) ([]byte, error) {
	return a.decryptBytes(ciphertext)
}
