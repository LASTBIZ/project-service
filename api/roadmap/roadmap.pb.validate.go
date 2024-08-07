// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/roadmap/roadmap.proto

package roadmap

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

// Validate checks the field values on CreateRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRoadmapRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoadmapRequestMultiError, or nil if none found.
func (m *CreateRoadmapRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoadmapRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetStartDate() == nil {
		err := CreateRoadmapRequestValidationError{
			field:  "StartDate",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEndDate() == nil {
		err := CreateRoadmapRequestValidationError{
			field:  "EndDate",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	// no validation rules for Name

	// no validation rules for Job

	// no validation rules for ProjectId

	if len(errors) > 0 {
		return CreateRoadmapRequestMultiError(errors)
	}

	return nil
}

// CreateRoadmapRequestMultiError is an error wrapping multiple validation
// errors returned by CreateRoadmapRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateRoadmapRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoadmapRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoadmapRequestMultiError) AllErrors() []error { return m }

// CreateRoadmapRequestValidationError is the validation error returned by
// CreateRoadmapRequest.Validate if the designated constraints aren't met.
type CreateRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoadmapRequestValidationError) ErrorName() string {
	return "CreateRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoadmapRequestValidationError{}

// Validate checks the field values on UpdateRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoadmapRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoadmapRequestMultiError, or nil if none found.
func (m *UpdateRoadmapRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoadmapRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetStartDate() == nil {
		err := UpdateRoadmapRequestValidationError{
			field:  "StartDate",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEndDate() == nil {
		err := UpdateRoadmapRequestValidationError{
			field:  "EndDate",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	// no validation rules for Name

	// no validation rules for Job

	// no validation rules for ProjectId

	// no validation rules for Id

	if len(errors) > 0 {
		return UpdateRoadmapRequestMultiError(errors)
	}

	return nil
}

// UpdateRoadmapRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateRoadmapRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateRoadmapRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoadmapRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoadmapRequestMultiError) AllErrors() []error { return m }

// UpdateRoadmapRequestValidationError is the validation error returned by
// UpdateRoadmapRequest.Validate if the designated constraints aren't met.
type UpdateRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoadmapRequestValidationError) ErrorName() string {
	return "UpdateRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoadmapRequestValidationError{}

// Validate checks the field values on DeleteRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRoadmapRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRoadmapRequestMultiError, or nil if none found.
func (m *DeleteRoadmapRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRoadmapRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteRoadmapRequestMultiError(errors)
	}

	return nil
}

// DeleteRoadmapRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteRoadmapRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteRoadmapRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRoadmapRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRoadmapRequestMultiError) AllErrors() []error { return m }

// DeleteRoadmapRequestValidationError is the validation error returned by
// DeleteRoadmapRequest.Validate if the designated constraints aren't met.
type DeleteRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRoadmapRequestValidationError) ErrorName() string {
	return "DeleteRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRoadmapRequestValidationError{}

// Validate checks the field values on ListRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListRoadmapRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRoadmapRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRoadmapRequestMultiError, or nil if none found.
func (m *ListRoadmapRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRoadmapRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageNum

	// no validation rules for PageSize

	// no validation rules for ProjectId

	if len(errors) > 0 {
		return ListRoadmapRequestMultiError(errors)
	}

	return nil
}

// ListRoadmapRequestMultiError is an error wrapping multiple validation errors
// returned by ListRoadmapRequest.ValidateAll() if the designated constraints
// aren't met.
type ListRoadmapRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRoadmapRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRoadmapRequestMultiError) AllErrors() []error { return m }

// ListRoadmapRequestValidationError is the validation error returned by
// ListRoadmapRequest.Validate if the designated constraints aren't met.
type ListRoadmapRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoadmapRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoadmapRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoadmapRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoadmapRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoadmapRequestValidationError) ErrorName() string {
	return "ListRoadmapRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRoadmapRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoadmapRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoadmapRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoadmapRequestValidationError{}

// Validate checks the field values on ListRoadmapReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListRoadmapReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRoadmapReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRoadmapReplyMultiError, or nil if none found.
func (m *ListRoadmapReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRoadmapReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRoadmaps() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListRoadmapReplyValidationError{
						field:  fmt.Sprintf("Roadmaps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListRoadmapReplyValidationError{
						field:  fmt.Sprintf("Roadmaps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRoadmapReplyValidationError{
					field:  fmt.Sprintf("Roadmaps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListRoadmapReplyMultiError(errors)
	}

	return nil
}

// ListRoadmapReplyMultiError is an error wrapping multiple validation errors
// returned by ListRoadmapReply.ValidateAll() if the designated constraints
// aren't met.
type ListRoadmapReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRoadmapReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRoadmapReplyMultiError) AllErrors() []error { return m }

// ListRoadmapReplyValidationError is the validation error returned by
// ListRoadmapReply.Validate if the designated constraints aren't met.
type ListRoadmapReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoadmapReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoadmapReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoadmapReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoadmapReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoadmapReplyValidationError) ErrorName() string { return "ListRoadmapReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListRoadmapReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoadmapReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoadmapReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoadmapReplyValidationError{}

// Validate checks the field values on ListRoadmapReply_Roadmap with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListRoadmapReply_Roadmap) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRoadmapReply_Roadmap with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRoadmapReply_RoadmapMultiError, or nil if none found.
func (m *ListRoadmapReply_Roadmap) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRoadmapReply_Roadmap) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStartDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListRoadmapReply_RoadmapValidationError{
					field:  "StartDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListRoadmapReply_RoadmapValidationError{
					field:  "StartDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListRoadmapReply_RoadmapValidationError{
				field:  "StartDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEndDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListRoadmapReply_RoadmapValidationError{
					field:  "EndDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListRoadmapReply_RoadmapValidationError{
					field:  "EndDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListRoadmapReply_RoadmapValidationError{
				field:  "EndDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Description

	// no validation rules for Name

	// no validation rules for Job

	// no validation rules for ProjectId

	// no validation rules for Id

	if len(errors) > 0 {
		return ListRoadmapReply_RoadmapMultiError(errors)
	}

	return nil
}

// ListRoadmapReply_RoadmapMultiError is an error wrapping multiple validation
// errors returned by ListRoadmapReply_Roadmap.ValidateAll() if the designated
// constraints aren't met.
type ListRoadmapReply_RoadmapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRoadmapReply_RoadmapMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRoadmapReply_RoadmapMultiError) AllErrors() []error { return m }

// ListRoadmapReply_RoadmapValidationError is the validation error returned by
// ListRoadmapReply_Roadmap.Validate if the designated constraints aren't met.
type ListRoadmapReply_RoadmapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoadmapReply_RoadmapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoadmapReply_RoadmapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoadmapReply_RoadmapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoadmapReply_RoadmapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoadmapReply_RoadmapValidationError) ErrorName() string {
	return "ListRoadmapReply_RoadmapValidationError"
}

// Error satisfies the builtin error interface
func (e ListRoadmapReply_RoadmapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoadmapReply_Roadmap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoadmapReply_RoadmapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoadmapReply_RoadmapValidationError{}
