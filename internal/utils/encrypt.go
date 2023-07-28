package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/Din-27/Go_job/internal/models"
	"golang.org/x/crypto/argon2"
)

func GenerateFromPassword(password string, p *models.Params) (encodedHash string, err error) {

	salt, err := generateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
