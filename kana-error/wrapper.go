package kana_error

import (
	"errors"
	"fmt"
)

var KanaError = errors.New("error of kana")

func TryErrorWrappingAndErrorsAs() {
	err := NewKanaCustom("lets generate err message", NewCustomErrorType("new Err"))

	var ce *CustomErrorType

	fmt.Println(errors.As(err, &ce))
}

func ReturnKanaError() error {
	return KanaError
}
func RunRootError() error {
	err := ReturnKanaError()

	newErr := fmt.Errorf("got a kana error: %w", err)

	return newErr
}
