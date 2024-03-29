// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package ioerror represents the interface "wasi:io/error@0.2.0-rc-2023-11-10".
package ioerror

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Error represents the resource "wasi:io/error@0.2.0-rc-2023-11-10#error".
//
// A resource which represents some error information.
//
// The only method provided by this resource is `to-debug-string`,
// which provides some human-readable information about the error.
//
// In the `wasi:io` package, this resource is returned through the
// `wasi:io/streams/stream-error` type.
//
// To provide more specific error information, other interfaces may
// provide functions to further "downcast" this error into more specific
// error information. For example, `error`s returned in streams derived
// from filesystem types to be described using the filesystem's own
// error-code type, using the function
// `wasi:filesystem/types/filesystem-error-code`, which takes a parameter
// `borrow<error>` and returns
// `option<wasi:filesystem/types/error-code>`.
//
// The set of functions which can "downcast" an `error` into a more
// concrete type is open.
//
//	resource error
type Error cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self Error) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:io/error@0.2.0-rc-2023-11-10 [resource-drop]error
//go:noescape
func (self Error) wasmimport_ResourceDrop()

// ToDebugString represents method "to-debug-string".
//
// Returns a string that is suitable to assist humans in debugging
// this error.
//
// WARNING: The returned string should not be consumed mechanically!
// It may change across platforms, hosts, or other implementation
// details. Parsing this string is a major platform-compatibility
// hazard.
//
//	to-debug-string: func() -> string
//
//go:nosplit
func (self Error) ToDebugString() string {
	var result string
	self.wasmimport_ToDebugString(&result)
	return result
}

//go:wasmimport wasi:io/error@0.2.0-rc-2023-11-10 [method]error.to-debug-string
//go:noescape
func (self Error) wasmimport_ToDebugString(result *string)
