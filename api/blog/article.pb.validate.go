// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: article.proto

package blog

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on CreateArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateArticleRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetTitle()) > 20 {
		return CreateArticleRequestValidationError{
			field:  "Title",
			reason: "value length must be at most 20 runes",
		}
	}

	if utf8.RuneCountInString(m.GetContent()) > 99999 {
		return CreateArticleRequestValidationError{
			field:  "Content",
			reason: "value length must be at most 99999 runes",
		}
	}

	// no validation rules for CreatedBy

	// no validation rules for UpdatedBy

	if val := m.GetType(); val < 1 || val > 100 {
		return CreateArticleRequestValidationError{
			field:  "Type",
			reason: "value must be inside range [1, 100]",
		}
	}

	// no validation rules for CityId
	if m.GetPublishedAt() == "" {
		return CreateArticleRequestValidationError{
			field:  "PublishedAt",
			reason: "value is required",
		}
	}

	return nil
}

// CreateArticleRequestValidationError is the validation error returned by
// CreateArticleRequest.Validate if the designated constraints aren't met.
type CreateArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleRequestValidationError) ErrorName() string {
	return "CreateArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleRequestValidationError{}

// Validate checks the field values on CreateArticleReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateArticleReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CreateArticleReplyValidationError is the validation error returned by
// CreateArticleReply.Validate if the designated constraints aren't met.
type CreateArticleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleReplyValidationError) ErrorName() string {
	return "CreateArticleReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleReplyValidationError{}

// Validate checks the field values on UpdateArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateArticleRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateArticleRequestValidationError is the validation error returned by
// UpdateArticleRequest.Validate if the designated constraints aren't met.
type UpdateArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateArticleRequestValidationError) ErrorName() string {
	return "UpdateArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateArticleRequestValidationError{}

// Validate checks the field values on UpdateArticleReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateArticleReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateArticleReplyValidationError is the validation error returned by
// UpdateArticleReply.Validate if the designated constraints aren't met.
type UpdateArticleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateArticleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateArticleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateArticleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateArticleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateArticleReplyValidationError) ErrorName() string {
	return "UpdateArticleReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateArticleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateArticleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateArticleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateArticleReplyValidationError{}

// Validate checks the field values on DeleteArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteArticleRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteArticleRequestValidationError is the validation error returned by
// DeleteArticleRequest.Validate if the designated constraints aren't met.
type DeleteArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteArticleRequestValidationError) ErrorName() string {
	return "DeleteArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteArticleRequestValidationError{}

// Validate checks the field values on DeleteArticleReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteArticleReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteArticleReplyValidationError is the validation error returned by
// DeleteArticleReply.Validate if the designated constraints aren't met.
type DeleteArticleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteArticleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteArticleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteArticleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteArticleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteArticleReplyValidationError) ErrorName() string {
	return "DeleteArticleReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteArticleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteArticleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteArticleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteArticleReplyValidationError{}

// Validate checks the field values on GetArticleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetArticleRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetArticleRequestValidationError is the validation error returned by
// GetArticleRequest.Validate if the designated constraints aren't met.
type GetArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleRequestValidationError) ErrorName() string {
	return "GetArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleRequestValidationError{}

// Validate checks the field values on GetArticleReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetArticleReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetArticleReplyValidationError is the validation error returned by
// GetArticleReply.Validate if the designated constraints aren't met.
type GetArticleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleReplyValidationError) ErrorName() string { return "GetArticleReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetArticleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleReplyValidationError{}

// Validate checks the field values on ListArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListArticleRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListArticleRequestValidationError is the validation error returned by
// ListArticleRequest.Validate if the designated constraints aren't met.
type ListArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListArticleRequestValidationError) ErrorName() string {
	return "ListArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListArticleRequestValidationError{}

// Validate checks the field values on ListArticleReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListArticleReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListArticleReplyValidationError is the validation error returned by
// ListArticleReply.Validate if the designated constraints aren't met.
type ListArticleReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListArticleReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListArticleReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListArticleReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListArticleReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListArticleReplyValidationError) ErrorName() string { return "ListArticleReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListArticleReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListArticleReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListArticleReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListArticleReplyValidationError{}
