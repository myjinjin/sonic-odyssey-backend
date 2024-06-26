package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type Encryptor interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(ciphertext string) (string, error)
}

type aesEncryptor struct {
	key []byte
}

func NewAESEncryptor(key string) (Encryptor, error) {
	if key == "" {
		return nil, errors.New("encryption key is not set")
	}
	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return &aesEncryptor{key: decodedKey}, nil
}

func (e *aesEncryptor) Encrypt(plaintext string) (string, error) {
	c, err := aes.NewCipher(e.key)
	if err != nil {
		return "", ErrCreatingCipher
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", ErrCreatingGCMCipher
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", ErrGeneratingNonce
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (e *aesEncryptor) Decrypt(ciphertext string) (string, error) {
	ct, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", ErrDecodingBase64Data
	}

	c, err := aes.NewCipher(e.key)
	if err != nil {
		return "", ErrCreatingCipher
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", ErrCreatingGCMCipher
	}

	nonceSize := gcm.NonceSize()
	if len(ct) < nonceSize {
		return "", ErrCipherTextTooShort
	}

	nonce, encryptedText := ct[:nonceSize], ct[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		return "", ErrDecryptingData
	}

	return string(plaintext), nil
}
