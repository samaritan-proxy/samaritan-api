// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/bootstrap/bootstrap.proto

package bootstrap

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Bootstrap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Bootstrap) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetInstance()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Instance",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLog()

		if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Log",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetStats()

		if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if m.GetAdmin() == nil {
		return BootstrapValidationError{
			field:  "Admin",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetAdmin()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Admin",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetStaticServices() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return BootstrapValidationError{
						field:  fmt.Sprintf("StaticServices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetDynamicSourceConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "DynamicSourceConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// BootstrapValidationError is the validation error returned by
// Bootstrap.Validate if the designated constraints aren't met.
type BootstrapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BootstrapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BootstrapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BootstrapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BootstrapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BootstrapValidationError) ErrorName() string { return "BootstrapValidationError" }

// Error satisfies the builtin error interface
func (e BootstrapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BootstrapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BootstrapValidationError{}

// Validate checks the field values on Sink with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Sink) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _Sink_Type_NotInLookup[m.GetType()]; ok {
		return SinkValidationError{
			field:  "Type",
			reason: "value must not be in list [0]",
		}
	}

	if len(m.GetEndpoint()) < 1 {
		return SinkValidationError{
			field:  "Endpoint",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// SinkValidationError is the validation error returned by Sink.Validate if the
// designated constraints aren't met.
type SinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SinkValidationError) ErrorName() string { return "SinkValidationError" }

// Error satisfies the builtin error interface
func (e SinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SinkValidationError{}

var _Sink_Type_NotInLookup = map[Sink_Type]struct{}{
	0: {},
}

// Validate checks the field values on Stats with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Stats) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSinks() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return StatsValidationError{
						field:  fmt.Sprintf("Sinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// StatsValidationError is the validation error returned by Stats.Validate if
// the designated constraints aren't met.
type StatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatsValidationError) ErrorName() string { return "StatsValidationError" }

// Error satisfies the builtin error interface
func (e StatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatsValidationError{}

// Validate checks the field values on Log with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Log) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Level

	{
		tmp := m.GetOutput()

		if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return LogValidationError{
					field:  "Output",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// LogValidationError is the validation error returned by Log.Validate if the
// designated constraints aren't met.
type LogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogValidationError) ErrorName() string { return "LogValidationError" }

// Error satisfies the builtin error interface
func (e LogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogValidationError{}

// Validate checks the field values on Admin with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Admin) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetBind() == nil {
		return AdminValidationError{
			field:  "Bind",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetBind()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return AdminValidationError{
					field:  "Bind",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// AdminValidationError is the validation error returned by Admin.Validate if
// the designated constraints aren't met.
type AdminValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminValidationError) ErrorName() string { return "AdminValidationError" }

// Error satisfies the builtin error interface
func (e AdminValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdmin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminValidationError{}

// Validate checks the field values on StaticService with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StaticService) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return StaticServiceValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	{
		tmp := m.GetConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return StaticServiceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if len(m.GetEndpoints()) < 1 {
		return StaticServiceValidationError{
			field:  "Endpoints",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return StaticServiceValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// StaticServiceValidationError is the validation error returned by
// StaticService.Validate if the designated constraints aren't met.
type StaticServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StaticServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StaticServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StaticServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StaticServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StaticServiceValidationError) ErrorName() string { return "StaticServiceValidationError" }

// Error satisfies the builtin error interface
func (e StaticServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStaticService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StaticServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StaticServiceValidationError{}

// Validate checks the field values on ConfigSource with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConfigSource) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetEndpoint()) < 1 {
		return ConfigSourceValidationError{
			field:  "Endpoint",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// ConfigSourceValidationError is the validation error returned by
// ConfigSource.Validate if the designated constraints aren't met.
type ConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigSourceValidationError) ErrorName() string { return "ConfigSourceValidationError" }

// Error satisfies the builtin error interface
func (e ConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigSourceValidationError{}

// Validate checks the field values on Log_Output with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Log_Output) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	// no validation rules for Target

	return nil
}

// Log_OutputValidationError is the validation error returned by
// Log_Output.Validate if the designated constraints aren't met.
type Log_OutputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Log_OutputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Log_OutputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Log_OutputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Log_OutputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Log_OutputValidationError) ErrorName() string { return "Log_OutputValidationError" }

// Error satisfies the builtin error interface
func (e Log_OutputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLog_Output.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Log_OutputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Log_OutputValidationError{}