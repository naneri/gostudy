package hasher

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func EncryptAndDecryptHmac() {
	key := "mysecretkey"
	message := "here is a message that I want to hash"
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	encrypted := h.Sum(nil)

	fmt.Println(encrypted)

	// did not finish this stuff
}
