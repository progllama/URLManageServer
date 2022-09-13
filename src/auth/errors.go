package auth

import (
	"errors"
	"net/http"
)

var (
	ErrSessionValueNotFound = errors.New("session value not found")
	ErrServerSideStateEmpty = errors.New("server side state is empty")
	ErrQueryStateEmpty      = errors.New("query state is empty")
	ErrStateNotMatch        = errors.New("state does not match")
	ErrOnExchange           = errors.New("get error on exchanging")
	ErrFetchUser            = errors.New("can't get user info")
	ErrAccessDatabase       = errors.New("db access error")
)

var (
	CodeServerSideStateEmpty = http.StatusBadRequest
	CodeQueryStateEmpty      = http.StatusBadRequest
	CodeStateNotMatch        = http.StatusBadRequest
	CodeOnExchange           = http.StatusInternalServerError
	CodeFetchUser            = http.StatusInternalServerError
	CodeAccessDatabase       = http.StatusInternalServerError
)

func ErrorCode(err error) int {
	switch err {
	case ErrServerSideStateEmpty:
		return CodeServerSideStateEmpty
	case ErrQueryStateEmpty:
		return CodeQueryStateEmpty
	case ErrStateNotMatch:
		return CodeStateNotMatch
	case ErrOnExchange:
		return CodeOnExchange
	case ErrFetchUser:
		return CodeFetchUser
	case ErrAccessDatabase:
		return CodeAccessDatabase
	default:
		return http.StatusInternalServerError
	}
}
