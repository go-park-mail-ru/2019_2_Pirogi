package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}