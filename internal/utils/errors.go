package utils

import (
	"errors"
	"fmt"
	"strings"
)

var (
	defaultSeparator = ";"
)

type ErrorGroupOption func(*ErrorGroup)

type ErrorGroup struct {
	name   string
	sep    string
	errors []error
}

func NewErrorGroup(options ...ErrorGroupOption) *ErrorGroup {
	eg := &ErrorGroup{
		sep: defaultSeparator,
	}

	for _, opt := range options {
		opt(eg)
	}

	return eg
}

func (eg *ErrorGroup) Join(err error) {
	if err == nil {
		return
	}

	eg.errors = append(eg.errors, err)
}

func (eg *ErrorGroup) Err() error {
	if len(eg.errors) == 0 {
		return nil
	}

	return errors.Join(eg.errors...)
}

func (eg *ErrorGroup) Error() string {
	errStrs := make([]string, 0, len(eg.errors))

	for _, err := range eg.errors {
		errStrs = append(errStrs, err.Error())
	}

	joinedErrs := strings.Join(errStrs, eg.sep)

	var prefix = eg.getPrefix()

	return fmt.Sprintf("[%s]:%s", prefix, joinedErrs)
}

func (eg *ErrorGroup) IsNil() bool {
	return len(eg.errors) == 0
}

func (eg *ErrorGroup) getPrefix() string {
	if eg.name == "" {
		return ""
	}

	return fmt.Sprintf("name=%s", eg.name)
}

func WithSeparator(sep string) ErrorGroupOption {
	return func(eg *ErrorGroup) {
		eg.sep = sep
	}
}

func WithName(name string) ErrorGroupOption {
	return func(eg *ErrorGroup) {
		eg.name = name
	}
}
