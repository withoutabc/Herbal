package util

import "errors"

var (
	TransactionError = errors.New("transaction error")
	ErrRowsAffected  = errors.New("rows affected error")
)

const (
	NoErrCode             = 0
	TransactionErrorCode  = 500
	InternalServerErrCode = 100
	ErrRecordNotFoundCode = 10000
	ErrRowsAffectedCode   = 10001
)
