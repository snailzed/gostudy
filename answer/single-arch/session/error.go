package session

import "github.com/pkg/errors"

var (
	ErrKeyNotExist     = errors.New("the key does not exists.")
	ErrSessionNotExist = errors.New("the sesion does not exists.")
)
