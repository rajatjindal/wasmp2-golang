// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package foo represents the interface "fermyon:spin/foo@2.0.0".
package foo

import "fmt"
// Greet represents function "fermyon:spin/foo@2.0.0#greet".
//
//	greet: func()
//
//go:nosplit
//export fermyon:spin/foo@2.0.0#greet
func Greet() {
	fmt.Println("hello greetings")
	wasmimport_Greet()
}

//go:wasmimport fermyon:spin/foo@2.0.0 greet
//go:noescape
func wasmimport_Greet()

type Interface interface {
	Greet()
}

var instance Interface
