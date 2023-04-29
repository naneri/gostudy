package hasher

import (
	"crypto/sha256"
	"fmt"
)

func Sha256HashWithNilSum() {
	src := []byte("Здесь могло быть написано, чем Go лучше Rust. " +
		"Но после хеширования уже не прочитаешь.")

	h := sha256.New()
	h.Write(src)
	dst := h.Sum(nil)

	fmt.Printf("%x\n", dst)

	fmt.Println("pause")
}

func Sha256HashWithouNilSum() {
	src := []byte("Здесь могло быть написано, чем Go лучше Rust. " +
		"Но после хеширования уже не прочитаешь.")

	h := sha256.New()
	h.Write(src)
	dst := h.Sum(nil)

	fmt.Printf("%x\n", dst)
}
