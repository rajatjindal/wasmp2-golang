// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package foo represents the interface "fermyon:spin/foo@2.0.0".
package foo

// Greet represents function "fermyon:spin/foo@2.0.0#greet".
//
//	greet: func() -> string
//
//go:nosplit
func Greet() string {
	var result string
	wasmimport_Greet(&result)
	return result
}

//go:wasmimport fermyon:spin/foo@2.0.0 greet
//go:noescape
func wasmimport_Greet(result *string)

type Interface interface {
	Greet(*string)
}

var instance Interface
