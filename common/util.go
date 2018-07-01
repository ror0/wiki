package common

import (
	"crypto/sha256"
)

func CalcHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return hash[:]
}

