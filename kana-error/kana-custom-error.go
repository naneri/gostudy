package kana_error

type KanaCustomError struct {
	message string
	err     error
}

func (kce *KanaCustomError) Error() string {
	return kce.message + ": " + kce.err.Error()
}

func (kce *KanaCustomError) Unwrap() error {
	return kce.err
}

func NewKanaCustom(message string, err error) error {
	return &KanaCustomError{
		message: message,
		err:     err,
	}
}
