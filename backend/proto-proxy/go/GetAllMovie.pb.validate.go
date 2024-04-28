// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: GetAllMovie.proto

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

// Validate checks the field values on GetAllMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAllMovieRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAllMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAllMovieRequestMultiError, or nil if none found.
func (m *GetAllMovieRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAllMovieRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Query

	if len(errors) > 0 {
		return GetAllMovieRequestMultiError(errors)
	}

	return nil
}

// GetAllMovieRequestMultiError is an error wrapping multiple validation errors
// returned by GetAllMovieRequest.ValidateAll() if the designated constraints
// aren't met.
type GetAllMovieRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAllMovieRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAllMovieRequestMultiError) AllErrors() []error { return m }

// GetAllMovieRequestValidationError is the validation error returned by
// GetAllMovieRequest.Validate if the designated constraints aren't met.
type GetAllMovieRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllMovieRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllMovieRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllMovieRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllMovieRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllMovieRequestValidationError) ErrorName() string {
	return "GetAllMovieRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAllMovieRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllMovieRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllMovieRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllMovieRequestValidationError{}

// Validate checks the field values on GetAllMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAllMovieResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAllMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAllMovieResponseMultiError, or nil if none found.
func (m *GetAllMovieResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAllMovieResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMovie() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetAllMovieResponseValidationError{
						field:  fmt.Sprintf("Movie[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetAllMovieResponseValidationError{
						field:  fmt.Sprintf("Movie[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAllMovieResponseValidationError{
					field:  fmt.Sprintf("Movie[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetAllMovieResponseMultiError(errors)
	}

	return nil
}

// GetAllMovieResponseMultiError is an error wrapping multiple validation
// errors returned by GetAllMovieResponse.ValidateAll() if the designated
// constraints aren't met.
type GetAllMovieResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAllMovieResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAllMovieResponseMultiError) AllErrors() []error { return m }

// GetAllMovieResponseValidationError is the validation error returned by
// GetAllMovieResponse.Validate if the designated constraints aren't met.
type GetAllMovieResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllMovieResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllMovieResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllMovieResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllMovieResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllMovieResponseValidationError) ErrorName() string {
	return "GetAllMovieResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAllMovieResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllMovieResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllMovieResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllMovieResponseValidationError{}
