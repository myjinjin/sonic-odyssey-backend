package hash

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type EmailHasher interface {
	HashEmail(email string) string
}

type sha256EmailHasher struct{}

var (
	onceEmail    sync.Once
	sha256Hasher EmailHasher
)

func SHA256EmailHasher() EmailHasher {
	onceEmail.Do(func() { sha256Hasher = &sha256EmailHasher{} })
	return sha256Hasher
}

func (h sha256EmailHasher) HashEmail(email string) string {
	hash := sha256.Sum256([]byte(email))
	return fmt.Sprintf("%x", hash)
}
