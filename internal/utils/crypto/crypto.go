package crypto

import (
	"sync"
)

var (
	defaultCrypto *AESCrypto
	once          sync.Once
)

// InitCrypto 从配置文件初始化默认加密实例
func InitCrypto() error {
	var err error
	once.Do(func() {
		defaultCrypto, err = NewAESCrypto()
	})
	return err
}

// Encrypt 使用默认实例加密字符串
func Encrypt(plaintext string) (string, error) {
	if defaultCrypto == nil {
		panic("crypto not initialized, call InitCrypto first")
	}
	return defaultCrypto.Encrypt(plaintext)
}

// Decrypt 使用默认实例解密字符串
func Decrypt(ciphertext string) (string, error) {
	if defaultCrypto == nil {
		panic("crypto not initialized, call InitCrypto first")
	}
	return defaultCrypto.Decrypt(ciphertext)
}

// EncryptBytes 使用默认实例加密字节数据
func EncryptBytes(data []byte) ([]byte, error) {
	if defaultCrypto == nil {
		panic("crypto not initialized, call InitCrypto first")
	}
	return defaultCrypto.EncryptBytes(data)
}

// DecryptBytes 使用默认实例解密字节数据
func DecryptBytes(ciphertext []byte) ([]byte, error) {
	if defaultCrypto == nil {
		panic("crypto not initialized, call InitCrypto first")
	}
	return defaultCrypto.DecryptBytes(ciphertext)
}
