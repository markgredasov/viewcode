package errs

import "errors"

var (
	ErrInvalidLanguage = errors.New("invalid language")
	ErrEmptyCode       = errors.New("empty code")
	ErrTooLargeCode    = errors.New("code is too large")
)
