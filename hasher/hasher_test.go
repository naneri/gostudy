package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestHasher(t *testing.T) {
	firstPart := "here is a start of the string "
	secondPart := "here is an end of the string"

	h := sha256.New()
	h.Write([]byte(firstPart))
	h.Write([]byte(secondPart))
	dst := h.Sum(nil)

	resultFirst := hex.EncodeToString(dst)

	byteArr := sha256.Sum256([]byte(firstPart + secondPart))
	secondResult := hex.EncodeToString(byteArr[:])

	if resultFirst != secondResult {
		t.Error("first and second hash are not equal")
	}

}
