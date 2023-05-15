package util

import "errors"

var (
	TransactionError = errors.New("transaction error")
	ErrRowsAffected  = errors.New("rows affected error")
	WrongPassword    = errors.New("wrong password")
)

const (
	NoErrCode             = 0
	TransactionErrorCode  = 500
	WrongPasswordCode     = 600
	InternalServerErrCode = 100
	ErrRecordNotFoundCode = 10000
	ErrRowsAffectedCode   = 10001
)
