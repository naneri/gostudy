package kana_error

type CustomErrorType struct {
	s string
}

func (ce *CustomErrorType) Error() string {
	return ce.s
}

func NewCustomErrorType(s string) error {
	return &CustomErrorType{s: s}
}
