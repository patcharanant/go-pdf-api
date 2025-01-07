package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")

	ErrCode = map[error]int{
		ErrInternalServerError: 500,
		ErrNotFound:            404,
		ErrConflict:            409,
		ErrBadParamInput:       400,
	}
)

func GetStatusCode(err error) int {
	code, ok := ErrCode[err]

	if !ok {
		return 500
	}
	return code
}
