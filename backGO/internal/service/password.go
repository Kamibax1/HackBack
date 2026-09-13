package service

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string, salt []byte) string {
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		3,
		64*1024,
		4,
		32,
	)

	return base64.StdEncoding.EncodeToString(hash)
}
