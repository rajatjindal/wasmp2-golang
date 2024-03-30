package inboundhttp

import (
	"fmt"
	"os"
)

// handler is the function that will be called by the http trigger in Spin.
var handler = defaultHandler

// defaultHandler is a placeholder for returning a useful error to stderr when
// the handler is not set.
var defaultHandler = func(req Request) *Response {
	fmt.Fprintln(os.Stderr, "http handler undefined")
	return &Response{
		Status: 400,
	}
}

// Handle sets the handler function for the http trigger.
// It must be set in an init() function.
func HandleFn(fn func(req Request) *Response) {
	wasmimport_HandleRequest = fn
}
