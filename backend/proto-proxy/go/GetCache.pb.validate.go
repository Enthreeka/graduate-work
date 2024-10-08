// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: GetCache.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on GetCacheRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCacheRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCacheRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCacheRequestMultiError, or nil if none found.
func (m *GetCacheRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCacheRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return GetCacheRequestMultiError(errors)
	}

	return nil
}

// GetCacheRequestMultiError is an error wrapping multiple validation errors
// returned by GetCacheRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCacheRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCacheRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCacheRequestMultiError) AllErrors() []error { return m }

// GetCacheRequestValidationError is the validation error returned by
// GetCacheRequest.Validate if the designated constraints aren't met.
type GetCacheRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCacheRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCacheRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCacheRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCacheRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCacheRequestValidationError) ErrorName() string { return "GetCacheRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetCacheRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCacheRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCacheRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCacheRequestValidationError{}

// Validate checks the field values on GetCacheResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCacheResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCacheResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCacheResponseMultiError, or nil if none found.
func (m *GetCacheResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCacheResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetCacheResponseMultiError(errors)
	}

	return nil
}

// GetCacheResponseMultiError is an error wrapping multiple validation errors
// returned by GetCacheResponse.ValidateAll() if the designated constraints
// aren't met.
type GetCacheResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCacheResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCacheResponseMultiError) AllErrors() []error { return m }

// GetCacheResponseValidationError is the validation error returned by
// GetCacheResponse.Validate if the designated constraints aren't met.
type GetCacheResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCacheResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCacheResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCacheResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCacheResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCacheResponseValidationError) ErrorName() string { return "GetCacheResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetCacheResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCacheResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCacheResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCacheResponseValidationError{}
