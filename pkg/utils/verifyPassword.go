package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

func VerifyPassword(password, encodedHash string) (bool, error) {
	// Split the stored hash into salt and hash
	parts := strings.Split(encodedHash, ".")
	if len(parts) != 2 {
		return false, errors.New("invalid stored hash format")
	}

	// Decode salt and stored hash from Base64
	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	storedHash, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	// Recompute the hash using the same salt
	computedHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, uint32(len(storedHash)))

	// Compare in constant time
	if subtle.ConstantTimeCompare(storedHash, computedHash) == 1 {
		return true, nil
	}

	return false, nil
}
