// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package variables represents the interface "fermyon:spin/variables@2.0.0".
package variables

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Error represents the variant "fermyon:spin/variables@2.0.0#error".
//
// The set of errors which may be raised by functions in this interface.
//
//	variant error {
//		invalid-name(string),
//		undefined(string),
//		provider(string),
//		other(string),
//	}
type Error cm.Variant[uint8, string, string]

// ErrorInvalidName returns a [Error] of case "invalid-name".
//
// The provided variable name is invalid.
func ErrorInvalidName(data string) Error {
	return cm.New[Error](0, data)
}

// InvalidName returns a non-nil *[string] if [Error] represents the variant case "invalid-name".
func (self *Error) InvalidName() *string {
	return cm.Case[string](self, 0)
}

// ErrorUndefined returns a [Error] of case "undefined".
//
// The provided variable is undefined.
func ErrorUndefined(data string) Error {
	return cm.New[Error](1, data)
}

// Undefined returns a non-nil *[string] if [Error] represents the variant case "undefined".
func (self *Error) Undefined() *string {
	return cm.Case[string](self, 1)
}

// ErrorProvider returns a [Error] of case "provider".
//
// A variables provider specific error has occurred.
func ErrorProvider(data string) Error {
	return cm.New[Error](2, data)
}

// Provider returns a non-nil *[string] if [Error] represents the variant case "provider".
func (self *Error) Provider() *string {
	return cm.Case[string](self, 2)
}

// ErrorOther returns a [Error] of case "other".
//
// Some implementation-specific error has occurred.
func ErrorOther(data string) Error {
	return cm.New[Error](3, data)
}

// Other returns a non-nil *[string] if [Error] represents the variant case "other".
func (self *Error) Other() *string {
	return cm.Case[string](self, 3)
}

// Get represents function "fermyon:spin/variables@2.0.0#get".
//
// Get an application variable value for the current component.
//
// The name must match one defined in in the component manifest.
//
//	get: func(name: string) -> result<string, error>
//
//go:nosplit
func Get(name string) cm.ErrResult[string, Error] {
	var result cm.ErrResult[string, Error]
	wasmimport_Get(name, &result)
	return result
}

//go:wasmimport fermyon:spin/variables@2.0.0 get
//go:noescape
func wasmimport_Get(name string, result *cm.ErrResult[string, Error])
