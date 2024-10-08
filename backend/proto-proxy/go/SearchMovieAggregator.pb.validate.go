// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: SearchMovieAggregator.proto

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

// Validate checks the field values on SearchMovieAggregatorRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SearchMovieAggregatorRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchMovieAggregatorRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SearchMovieAggregatorRequestMultiError, or nil if none found.
func (m *SearchMovieAggregatorRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchMovieAggregatorRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RedisKey

	// no validation rules for Query

	if len(errors) > 0 {
		return SearchMovieAggregatorRequestMultiError(errors)
	}

	return nil
}

// SearchMovieAggregatorRequestMultiError is an error wrapping multiple
// validation errors returned by SearchMovieAggregatorRequest.ValidateAll() if
// the designated constraints aren't met.
type SearchMovieAggregatorRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchMovieAggregatorRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchMovieAggregatorRequestMultiError) AllErrors() []error { return m }

// SearchMovieAggregatorRequestValidationError is the validation error returned
// by SearchMovieAggregatorRequest.Validate if the designated constraints
// aren't met.
type SearchMovieAggregatorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchMovieAggregatorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchMovieAggregatorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchMovieAggregatorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchMovieAggregatorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchMovieAggregatorRequestValidationError) ErrorName() string {
	return "SearchMovieAggregatorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchMovieAggregatorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchMovieAggregatorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchMovieAggregatorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchMovieAggregatorRequestValidationError{}

// Validate checks the field values on SearchMovieAggregatorResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SearchMovieAggregatorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchMovieAggregatorResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// SearchMovieAggregatorResponseMultiError, or nil if none found.
func (m *SearchMovieAggregatorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchMovieAggregatorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHits()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SearchMovieAggregatorResponseValidationError{
					field:  "Hits",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SearchMovieAggregatorResponseValidationError{
					field:  "Hits",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHits()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SearchMovieAggregatorResponseValidationError{
				field:  "Hits",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SearchMovieAggregatorResponseMultiError(errors)
	}

	return nil
}

// SearchMovieAggregatorResponseMultiError is an error wrapping multiple
// validation errors returned by SearchMovieAggregatorResponse.ValidateAll()
// if the designated constraints aren't met.
type SearchMovieAggregatorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchMovieAggregatorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchMovieAggregatorResponseMultiError) AllErrors() []error { return m }

// SearchMovieAggregatorResponseValidationError is the validation error
// returned by SearchMovieAggregatorResponse.Validate if the designated
// constraints aren't met.
type SearchMovieAggregatorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchMovieAggregatorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchMovieAggregatorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchMovieAggregatorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchMovieAggregatorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchMovieAggregatorResponseValidationError) ErrorName() string {
	return "SearchMovieAggregatorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SearchMovieAggregatorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchMovieAggregatorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchMovieAggregatorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchMovieAggregatorResponseValidationError{}
