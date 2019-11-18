package common

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

