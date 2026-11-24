package errors

import "errors"

var ErrAlreadyExists error = errors.New("already exists")
var ErrPhoneNumUnknown error = errors.New("unknown number")
