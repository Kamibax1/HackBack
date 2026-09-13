package service

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func BuildGameHash(
	sessionID string,
	score int,
	secret string,
) string {

	data := sessionID +
		strconv.Itoa(score) +
		secret

	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}
