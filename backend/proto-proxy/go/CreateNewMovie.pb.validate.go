// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: CreateNewMovie.proto

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

// Validate checks the field values on CreateNewMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNewMovieRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNewMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNewMovieRequestMultiError, or nil if none found.
func (m *CreateNewMovieRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNewMovieRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMovie()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateNewMovieRequestValidationError{
					field:  "Movie",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateNewMovieRequestValidationError{
					field:  "Movie",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMovie()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateNewMovieRequestValidationError{
				field:  "Movie",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateNewMovieRequestMultiError(errors)
	}

	return nil
}

// CreateNewMovieRequestMultiError is an error wrapping multiple validation
// errors returned by CreateNewMovieRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateNewMovieRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNewMovieRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNewMovieRequestMultiError) AllErrors() []error { return m }

// CreateNewMovieRequestValidationError is the validation error returned by
// CreateNewMovieRequest.Validate if the designated constraints aren't met.
type CreateNewMovieRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNewMovieRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNewMovieRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNewMovieRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNewMovieRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNewMovieRequestValidationError) ErrorName() string {
	return "CreateNewMovieRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNewMovieRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNewMovieRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNewMovieRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNewMovieRequestValidationError{}

// Validate checks the field values on CreateNewMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNewMovieResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNewMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNewMovieResponseMultiError, or nil if none found.
func (m *CreateNewMovieResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNewMovieResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateNewMovieResponseMultiError(errors)
	}

	return nil
}

// CreateNewMovieResponseMultiError is an error wrapping multiple validation
// errors returned by CreateNewMovieResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateNewMovieResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNewMovieResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNewMovieResponseMultiError) AllErrors() []error { return m }

// CreateNewMovieResponseValidationError is the validation error returned by
// CreateNewMovieResponse.Validate if the designated constraints aren't met.
type CreateNewMovieResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNewMovieResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNewMovieResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNewMovieResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNewMovieResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNewMovieResponseValidationError) ErrorName() string {
	return "CreateNewMovieResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNewMovieResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNewMovieResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNewMovieResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNewMovieResponseValidationError{}
