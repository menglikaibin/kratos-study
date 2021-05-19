package biz

import (
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/genproto/googleapis/api/error_reason"
	"net/http"
)

type Error struct {
	Code     int
	Domain   string
	Reason   string
	Message  string
	Metadata map[string]string
}

func (e *Error) WithMetaData(md map[string]string) *Error {
	err := *e
	err.Metadata = md

	return &err
}

func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

func NewError(code int, domain, reason, message string) *Error {
	return &Error{Code: code, Domain: domain, Reason: reason, Message: message}
}

func NewAccessTokenExpired(format string, a ...interface{}) *Error {
	return NewError(http.StatusBadRequest, "user", error_reason.ErrorReason_ACCESS_TOKEN_EXPIRED.String(), fmt.Sprintf(format, a...))
}

func IsAccessTokenExpired(err error) bool {
	if e := new(Error); errors.As(err, &e) {
		return e.Domain == "user" && e.Reason == error_reason.ErrorReason_ACCESS_TOKEN_EXPIRED.String()
	}
	return false
}

var (
	ErrAccessTokenExpired = NewError(http.StatusBadRequest, "user", error_reason.ErrorReason_ACCESS_TOKEN_EXPIRED.String(), "access token expired")
)





