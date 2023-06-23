package adb

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/exidler/goadb/internal/errors"
)

var (
	whitespaceRegex = regexp.MustCompile(`^\s*$`)
)

func containsWhitespace(str string) bool {
	return strings.ContainsAny(str, " \t\v")
}

func isBlank(str string) bool {
	return whitespaceRegex.MatchString(str)
}

func wrapClientError(err error, client interface{}, operation string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*errors.Err); !ok {
		panic("err is not a *Err: " + err.Error())
	}

	clientType := reflect.TypeOf(client)

	return &errors.Err{
		Code:    err.(*errors.Err).Code,
		Cause:   err,
		Message: fmt.Sprintf("error performing %s on %s", fmt.Sprintf(operation, args...), clientType),
		Details: client,
	}
}

/*
WrapErrf returns an *Err that wraps another *Err and has the same ErrCode.
Panics if cause is not an *Err.

To wrap generic errors, use WrapErrorf.
*/
func WrapErrf(cause error, format string, args ...interface{}) error {
	return errors.WrapErrf(cause, format, args)
}

func Errorf(code ErrCode, format string, args ...interface{}) error {
	return errors.Errorf(errors.ErrCode(code), format, args)
}
