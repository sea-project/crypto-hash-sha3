package crypto_hash_sha3

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"testing"
)

func Test_sha3(t *testing.T) {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte("123"))
	t.Log(hex.EncodeToString(hasher.Sum(nil)))
	t.Log(hex.EncodeToString(Keccak256([]byte("123"))))
}
