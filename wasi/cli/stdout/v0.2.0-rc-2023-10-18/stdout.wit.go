// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package stdout represents the interface "wasi:cli/stdout@0.2.0-rc-2023-10-18".
package stdout

import (
	streams "github.com/rajatjindal/wasip2-golang/wasi/io/streams/v0.2.0-rc-2023-10-18"
)

// OutputStream represents the resource "wasi:io/streams@0.2.0-rc-2023-10-18#output-stream".
//
// See [streams.OutputStream] for more information.
type OutputStream = streams.OutputStream

// GetStdout represents function "wasi:cli/stdout@0.2.0-rc-2023-10-18#get-stdout".
//
//	get-stdout: func() -> own<output-stream>
//
//go:nosplit
func GetStdout() OutputStream {
	return wasmimport_GetStdout()
}

//go:wasmimport wasi:cli/stdout@0.2.0-rc-2023-10-18 get-stdout
//go:noescape
func wasmimport_GetStdout() OutputStream
