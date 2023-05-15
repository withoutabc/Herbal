package util

import "errors"

var (
	TransactionError = errors.New("transaction error")
)

const (
	NoErrCode             = 0
	TransactionErrorCode  = 500
	InternalServerErrCode = 100
	ErrRecordNotFoundCode = 10000
)
