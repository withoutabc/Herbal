package errutil

import "herbalBody/anotherone/util2/codes"

type CodeError struct {
	Code   int
	reason string
}

func (e CodeError) Error() string {
	return e.reason
}

func NewWithCode(code int) error {
	return CodeError{
		Code:   code,
		reason: codes.CodeErrorMap[code],
	}
}

func ToCodeError(code int, err error) error {
	if err == nil {
		return nil
	}
	return CodeError{
		Code:   code,
		reason: codes.CodeErrorMap[code],
	}
}
