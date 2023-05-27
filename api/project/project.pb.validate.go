// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/project/project.proto

package project

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

// Validate checks the field values on CreateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProjectRequestMultiError, or nil if none found.
func (m *CreateProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !strings.HasPrefix(m.GetRoadmapImgUrl(), "https://") {
		err := CreateProjectRequestValidationError{
			field:  "RoadmapImgUrl",
			reason: "value does not have prefix \"https://\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !strings.HasPrefix(m.GetMainImageUrl(), "https://") {
		err := CreateProjectRequestValidationError{
			field:  "MainImageUrl",
			reason: "value does not have prefix \"https://\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStartDate() == nil {
		err := CreateProjectRequestValidationError{
			field:  "StartDate",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEndDate() == nil {
		err := CreateProjectRequestValidationError{
			field:  "EndDate",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetLocation()) < 1 {
		err := CreateProjectRequestValidationError{
			field:  "Location",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CreateProjectRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := CreateProjectRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCategoryId() <= 0 {
		err := CreateProjectRequestValidationError{
			field:  "CategoryId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetNeedBudget() <= 0 {
		err := CreateProjectRequestValidationError{
			field:  "NeedBudget",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateProjectRequestMultiError(errors)
	}

	return nil
}

// CreateProjectRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProjectRequestMultiError) AllErrors() []error { return m }

// CreateProjectRequestValidationError is the validation error returned by
// CreateProjectRequest.Validate if the designated constraints aren't met.
type CreateProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProjectRequestValidationError) ErrorName() string {
	return "CreateProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProjectRequestValidationError{}

// Validate checks the field values on UpdateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProjectRequestMultiError, or nil if none found.
func (m *UpdateProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !strings.HasPrefix(m.GetRoadmapImgUrl(), "https://") {
		err := UpdateProjectRequestValidationError{
			field:  "RoadmapImgUrl",
			reason: "value does not have prefix \"https://\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !strings.HasPrefix(m.GetMainImageUrl(), "https://") {
		err := UpdateProjectRequestValidationError{
			field:  "MainImageUrl",
			reason: "value does not have prefix \"https://\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStartDate() == nil {
		err := UpdateProjectRequestValidationError{
			field:  "StartDate",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEndDate() == nil {
		err := UpdateProjectRequestValidationError{
			field:  "EndDate",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetLocation()) < 1 {
		err := UpdateProjectRequestValidationError{
			field:  "Location",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UpdateProjectRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := UpdateProjectRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCategoryId() <= 0 {
		err := UpdateProjectRequestValidationError{
			field:  "CategoryId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetNeedBudget() <= 0 {
		err := UpdateProjectRequestValidationError{
			field:  "NeedBudget",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetId() <= 0 {
		err := UpdateProjectRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateProjectRequestMultiError(errors)
	}

	return nil
}

// UpdateProjectRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProjectRequestMultiError) AllErrors() []error { return m }

// UpdateProjectRequestValidationError is the validation error returned by
// UpdateProjectRequest.Validate if the designated constraints aren't met.
type UpdateProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProjectRequestValidationError) ErrorName() string {
	return "UpdateProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProjectRequestValidationError{}

// Validate checks the field values on DeleteProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProjectRequestMultiError, or nil if none found.
func (m *DeleteProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteProjectRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteProjectRequestMultiError(errors)
	}

	return nil
}

// DeleteProjectRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProjectRequestMultiError) AllErrors() []error { return m }

// DeleteProjectRequestValidationError is the validation error returned by
// DeleteProjectRequest.Validate if the designated constraints aren't met.
type DeleteProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProjectRequestValidationError) ErrorName() string {
	return "DeleteProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProjectRequestValidationError{}

// Validate checks the field values on GetProjectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProjectRequestMultiError, or nil if none found.
func (m *GetProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetProjectRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetProjectRequestMultiError(errors)
	}

	return nil
}

// GetProjectRequestMultiError is an error wrapping multiple validation errors
// returned by GetProjectRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProjectRequestMultiError) AllErrors() []error { return m }

// GetProjectRequestValidationError is the validation error returned by
// GetProjectRequest.Validate if the designated constraints aren't met.
type GetProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectRequestValidationError) ErrorName() string {
	return "GetProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectRequestValidationError{}

// Validate checks the field values on GetProjectReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProjectReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProjectReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProjectReplyMultiError, or nil if none found.
func (m *GetProjectReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProjectReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoadmapImgUrl

	// no validation rules for MainImageUrl

	if all {
		switch v := interface{}(m.GetStartDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProjectReplyValidationError{
					field:  "StartDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProjectReplyValidationError{
					field:  "StartDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProjectReplyValidationError{
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
				errors = append(errors, GetProjectReplyValidationError{
					field:  "EndDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProjectReplyValidationError{
					field:  "EndDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProjectReplyValidationError{
				field:  "EndDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Location

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for CategoryId

	// no validation rules for NeedBudget

	// no validation rules for Id

	// no validation rules for CurrentBudget

	if len(errors) > 0 {
		return GetProjectReplyMultiError(errors)
	}

	return nil
}

// GetProjectReplyMultiError is an error wrapping multiple validation errors
// returned by GetProjectReply.ValidateAll() if the designated constraints
// aren't met.
type GetProjectReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProjectReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProjectReplyMultiError) AllErrors() []error { return m }

// GetProjectReplyValidationError is the validation error returned by
// GetProjectReply.Validate if the designated constraints aren't met.
type GetProjectReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectReplyValidationError) ErrorName() string { return "GetProjectReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetProjectReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectReplyValidationError{}

// Validate checks the field values on ListProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProjectRequestMultiError, or nil if none found.
func (m *ListProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Keywords

	// no validation rules for PageNum

	// no validation rules for PageSize

	// no validation rules for CategoryId

	if len(errors) > 0 {
		return ListProjectRequestMultiError(errors)
	}

	return nil
}

// ListProjectRequestMultiError is an error wrapping multiple validation errors
// returned by ListProjectRequest.ValidateAll() if the designated constraints
// aren't met.
type ListProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProjectRequestMultiError) AllErrors() []error { return m }

// ListProjectRequestValidationError is the validation error returned by
// ListProjectRequest.Validate if the designated constraints aren't met.
type ListProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectRequestValidationError) ErrorName() string {
	return "ListProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectRequestValidationError{}

// Validate checks the field values on ListProjectReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListProjectReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProjectReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProjectReplyMultiError, or nil if none found.
func (m *ListProjectReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProjectReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListProjectReplyValidationError{
						field:  fmt.Sprintf("Projects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListProjectReplyValidationError{
						field:  fmt.Sprintf("Projects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProjectReplyValidationError{
					field:  fmt.Sprintf("Projects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListProjectReplyMultiError(errors)
	}

	return nil
}

// ListProjectReplyMultiError is an error wrapping multiple validation errors
// returned by ListProjectReply.ValidateAll() if the designated constraints
// aren't met.
type ListProjectReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProjectReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProjectReplyMultiError) AllErrors() []error { return m }

// ListProjectReplyValidationError is the validation error returned by
// ListProjectReply.Validate if the designated constraints aren't met.
type ListProjectReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectReplyValidationError) ErrorName() string { return "ListProjectReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListProjectReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectReplyValidationError{}

// Validate checks the field values on ProjectInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProjectInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProjectInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProjectInfoResponseMultiError, or nil if none found.
func (m *ProjectInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProjectInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MainImageUrl

	if all {
		switch v := interface{}(m.GetStartDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectInfoResponseValidationError{
					field:  "StartDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectInfoResponseValidationError{
					field:  "StartDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectInfoResponseValidationError{
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
				errors = append(errors, ProjectInfoResponseValidationError{
					field:  "EndDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectInfoResponseValidationError{
					field:  "EndDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectInfoResponseValidationError{
				field:  "EndDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Location

	// no validation rules for Name

	// no validation rules for CategoryId

	// no validation rules for Id

	if len(errors) > 0 {
		return ProjectInfoResponseMultiError(errors)
	}

	return nil
}

// ProjectInfoResponseMultiError is an error wrapping multiple validation
// errors returned by ProjectInfoResponse.ValidateAll() if the designated
// constraints aren't met.
type ProjectInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProjectInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProjectInfoResponseMultiError) AllErrors() []error { return m }

// ProjectInfoResponseValidationError is the validation error returned by
// ProjectInfoResponse.Validate if the designated constraints aren't met.
type ProjectInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectInfoResponseValidationError) ErrorName() string {
	return "ProjectInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectInfoResponseValidationError{}

// Validate checks the field values on InvestProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InvestProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InvestProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InvestProjectRequestMultiError, or nil if none found.
func (m *InvestProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InvestProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := InvestProjectRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetInvestorId() <= 0 {
		err := InvestProjectRequestValidationError{
			field:  "InvestorId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMoney() <= 0 {
		err := InvestProjectRequestValidationError{
			field:  "Money",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return InvestProjectRequestMultiError(errors)
	}

	return nil
}

// InvestProjectRequestMultiError is an error wrapping multiple validation
// errors returned by InvestProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type InvestProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InvestProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InvestProjectRequestMultiError) AllErrors() []error { return m }

// InvestProjectRequestValidationError is the validation error returned by
// InvestProjectRequest.Validate if the designated constraints aren't met.
type InvestProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InvestProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InvestProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InvestProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InvestProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InvestProjectRequestValidationError) ErrorName() string {
	return "InvestProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InvestProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInvestProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InvestProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InvestProjectRequestValidationError{}