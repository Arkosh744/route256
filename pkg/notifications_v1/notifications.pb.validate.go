// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: notifications.proto

package notifications_v1

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

// Validate checks the field values on ListUserHistoryDayRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUserHistoryDayRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserHistoryDayRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserHistoryDayRequestMultiError, or nil if none found.
func (m *ListUserHistoryDayRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserHistoryDayRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUser() <= 0 {
		err := ListUserHistoryDayRequestValidationError{
			field:  "User",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListUserHistoryDayRequestMultiError(errors)
	}

	return nil
}

// ListUserHistoryDayRequestMultiError is an error wrapping multiple validation
// errors returned by ListUserHistoryDayRequest.ValidateAll() if the
// designated constraints aren't met.
type ListUserHistoryDayRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserHistoryDayRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserHistoryDayRequestMultiError) AllErrors() []error { return m }

// ListUserHistoryDayRequestValidationError is the validation error returned by
// ListUserHistoryDayRequest.Validate if the designated constraints aren't met.
type ListUserHistoryDayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserHistoryDayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserHistoryDayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserHistoryDayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserHistoryDayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserHistoryDayRequestValidationError) ErrorName() string {
	return "ListUserHistoryDayRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListUserHistoryDayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserHistoryDayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserHistoryDayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserHistoryDayRequestValidationError{}

// Validate checks the field values on ListUserHistoryDayResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUserHistoryDayResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserHistoryDayResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserHistoryDayResponseMultiError, or nil if none found.
func (m *ListUserHistoryDayResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserHistoryDayResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUserHistoryDayResponseValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUserHistoryDayResponseValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserHistoryDayResponseValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListUserHistoryDayResponseMultiError(errors)
	}

	return nil
}

// ListUserHistoryDayResponseMultiError is an error wrapping multiple
// validation errors returned by ListUserHistoryDayResponse.ValidateAll() if
// the designated constraints aren't met.
type ListUserHistoryDayResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserHistoryDayResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserHistoryDayResponseMultiError) AllErrors() []error { return m }

// ListUserHistoryDayResponseValidationError is the validation error returned
// by ListUserHistoryDayResponse.Validate if the designated constraints aren't met.
type ListUserHistoryDayResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserHistoryDayResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserHistoryDayResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserHistoryDayResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserHistoryDayResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserHistoryDayResponseValidationError) ErrorName() string {
	return "ListUserHistoryDayResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListUserHistoryDayResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserHistoryDayResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserHistoryDayResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserHistoryDayResponseValidationError{}

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Message) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MessageMultiError, or nil if none found.
func (m *Message) ValidateAll() error {
	return m.validate(true)
}

func (m *Message) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for Status

	if len(errors) > 0 {
		return MessageMultiError(errors)
	}

	return nil
}

// MessageMultiError is an error wrapping multiple validation errors returned
// by Message.ValidateAll() if the designated constraints aren't met.
type MessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageMultiError) AllErrors() []error { return m }

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}