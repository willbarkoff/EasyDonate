package util

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomBytes generates a cryptographically-secure pseudorandom slice of bytes of length n.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString generates a cryptographically-secure pseudorandom string of length s
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return (base64.URLEncoding.EncodeToString(b))[:s], err
}
