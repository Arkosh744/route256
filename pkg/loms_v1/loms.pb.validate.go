// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: loms.proto

package loms_v1

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

// Validate checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderRequestMultiError, or nil if none found.
func (m *CreateOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUser() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "User",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateOrderRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderRequestValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateOrderRequestMultiError(errors)
	}

	return nil
}

// CreateOrderRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderRequestMultiError) AllErrors() []error { return m }

// CreateOrderRequestValidationError is the validation error returned by
// CreateOrderRequest.Validate if the designated constraints aren't met.
type CreateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequestValidationError) ErrorName() string {
	return "CreateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequestValidationError{}

// Validate checks the field values on Item with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Item with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ItemMultiError, or nil if none found.
func (m *Item) ValidateAll() error {
	return m.validate(true)
}

func (m *Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sku

	if val := m.GetCount(); val <= 0 || val >= 65535 {
		err := ItemValidationError{
			field:  "Count",
			reason: "value must be inside range (0, 65535)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ItemMultiError(errors)
	}

	return nil
}

// ItemMultiError is an error wrapping multiple validation errors returned by
// Item.ValidateAll() if the designated constraints aren't met.
type ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemMultiError) AllErrors() []error { return m }

// ItemValidationError is the validation error returned by Item.Validate if the
// designated constraints aren't met.
type ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemValidationError) ErrorName() string { return "ItemValidationError" }

// Error satisfies the builtin error interface
func (e ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemValidationError{}

// Validate checks the field values on CreateOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderResponseMultiError, or nil if none found.
func (m *CreateOrderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return CreateOrderResponseMultiError(errors)
	}

	return nil
}

// CreateOrderResponseMultiError is an error wrapping multiple validation
// errors returned by CreateOrderResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateOrderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderResponseMultiError) AllErrors() []error { return m }

// CreateOrderResponseValidationError is the validation error returned by
// CreateOrderResponse.Validate if the designated constraints aren't met.
type CreateOrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderResponseValidationError) ErrorName() string {
	return "CreateOrderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderResponseValidationError{}

// Validate checks the field values on ListOrderResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListOrderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListOrderResponseMultiError, or nil if none found.
func (m *ListOrderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for User

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListOrderResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListOrderResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListOrderResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListOrderResponseMultiError(errors)
	}

	return nil
}

// ListOrderResponseMultiError is an error wrapping multiple validation errors
// returned by ListOrderResponse.ValidateAll() if the designated constraints
// aren't met.
type ListOrderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrderResponseMultiError) AllErrors() []error { return m }

// ListOrderResponseValidationError is the validation error returned by
// ListOrderResponse.Validate if the designated constraints aren't met.
type ListOrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderResponseValidationError) ErrorName() string {
	return "ListOrderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListOrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderResponseValidationError{}

// Validate checks the field values on StocksRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StocksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StocksRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StocksRequestMultiError, or
// nil if none found.
func (m *StocksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StocksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sku

	if len(errors) > 0 {
		return StocksRequestMultiError(errors)
	}

	return nil
}

// StocksRequestMultiError is an error wrapping multiple validation errors
// returned by StocksRequest.ValidateAll() if the designated constraints
// aren't met.
type StocksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StocksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StocksRequestMultiError) AllErrors() []error { return m }

// StocksRequestValidationError is the validation error returned by
// StocksRequest.Validate if the designated constraints aren't met.
type StocksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StocksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StocksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StocksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StocksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StocksRequestValidationError) ErrorName() string { return "StocksRequestValidationError" }

// Error satisfies the builtin error interface
func (e StocksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStocksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StocksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StocksRequestValidationError{}

// Validate checks the field values on StocksResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StocksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StocksResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StocksResponseMultiError,
// or nil if none found.
func (m *StocksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StocksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetStocks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StocksResponseValidationError{
						field:  fmt.Sprintf("Stocks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StocksResponseValidationError{
						field:  fmt.Sprintf("Stocks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StocksResponseValidationError{
					field:  fmt.Sprintf("Stocks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return StocksResponseMultiError(errors)
	}

	return nil
}

// StocksResponseMultiError is an error wrapping multiple validation errors
// returned by StocksResponse.ValidateAll() if the designated constraints
// aren't met.
type StocksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StocksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StocksResponseMultiError) AllErrors() []error { return m }

// StocksResponseValidationError is the validation error returned by
// StocksResponse.Validate if the designated constraints aren't met.
type StocksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StocksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StocksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StocksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StocksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StocksResponseValidationError) ErrorName() string { return "StocksResponseValidationError" }

// Error satisfies the builtin error interface
func (e StocksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStocksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StocksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StocksResponseValidationError{}

// Validate checks the field values on Stock with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Stock) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Stock with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StockMultiError, or nil if none found.
func (m *Stock) ValidateAll() error {
	return m.validate(true)
}

func (m *Stock) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WarehouseId

	// no validation rules for Count

	if len(errors) > 0 {
		return StockMultiError(errors)
	}

	return nil
}

// StockMultiError is an error wrapping multiple validation errors returned by
// Stock.ValidateAll() if the designated constraints aren't met.
type StockMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockMultiError) AllErrors() []error { return m }

// StockValidationError is the validation error returned by Stock.Validate if
// the designated constraints aren't met.
type StockValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockValidationError) ErrorName() string { return "StockValidationError" }

// Error satisfies the builtin error interface
func (e StockValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStock.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockValidationError{}

// Validate checks the field values on OrderIDRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderIDRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderIDRequestMultiError,
// or nil if none found.
func (m *OrderIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return OrderIDRequestMultiError(errors)
	}

	return nil
}

// OrderIDRequestMultiError is an error wrapping multiple validation errors
// returned by OrderIDRequest.ValidateAll() if the designated constraints
// aren't met.
type OrderIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderIDRequestMultiError) AllErrors() []error { return m }

// OrderIDRequestValidationError is the validation error returned by
// OrderIDRequest.Validate if the designated constraints aren't met.
type OrderIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderIDRequestValidationError) ErrorName() string { return "OrderIDRequestValidationError" }

// Error satisfies the builtin error interface
func (e OrderIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderIDRequestValidationError{}
