// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package run represents the interface "wasi:cli/run@0.2.0-rc-2023-11-10".
package run

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Run represents function "wasi:cli/run@0.2.0-rc-2023-11-10#run".
//
// Run the program.
//
//	run: func() -> result
//
//go:nosplit
func Run() cm.Result {
	return wasmimport_Run()
}

//go:wasmimport wasi:cli/run@0.2.0-rc-2023-11-10 run
//go:noescape
func wasmimport_Run() cm.Result
