package errutil

import "herbalBody/anotherone/util2/codes"

type CodeError struct {
	Code   int
	reason string
}

func (e CodeError) Error() string {
	return e.reason
}

func NewWithCode(code int) CodeError {
	return CodeError{
		Code:   code,
		reason: codes.CodeErrorMap[code],
	}
}
