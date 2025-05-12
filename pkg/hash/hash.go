package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func CreateHash(data any) (string, error) {
	str, ok := data.(string)
	if !ok {
		return "", errors.New("cannot to convert data to string")
	}
	hash := sha256.Sum256([]byte(str))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr, nil
}
