package hash

import (
	"crypto/sha1"
	"fmt"
)

// PasswordHasher - Интерфейс для взаимодействия с хэшером
type PasswordHasher interface {
	Hash(password string) string
}

// SHA1Hasher - структура для работы с sha1 хэшером
type SHA1Hasher struct {
	salt string
}

// NewSHA1Hasher - Создание нового хэшера
func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{
		salt: salt,
	}
}

// Hash - хэширует полученный пароль с солью
func (h *SHA1Hasher) Hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt)))
}
