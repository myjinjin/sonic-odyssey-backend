package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

type EmailHasher interface {
	HashEmail(email string) string
}

type sha256EmailHasher struct{}

func SHA256EmailHasher() EmailHasher {
	return &sha256EmailHasher{}
}

func (h sha256EmailHasher) HashEmail(email string) string {
	hash := sha256.Sum256([]byte(email))
	return hex.EncodeToString(hash[:])
}
