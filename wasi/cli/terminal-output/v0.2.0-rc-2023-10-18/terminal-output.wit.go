// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package terminaloutput represents the interface "wasi:cli/terminal-output@0.2.0-rc-2023-10-18".
package terminaloutput

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// TerminalOutput represents the resource "wasi:cli/terminal-output@0.2.0-rc-2023-10-18#terminal-output".
//
// The output side of a terminal.
//
//	resource terminal-output
type TerminalOutput cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self TerminalOutput) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:cli/terminal-output@0.2.0-rc-2023-10-18 [resource-drop]terminal-output
//go:noescape
func (self TerminalOutput) wasmimport_ResourceDrop()
