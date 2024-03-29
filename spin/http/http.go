package http

import (
	"fmt"
	"net/http"
	"os"

	inboundhttp "github.com/rajatjindal/wasip2-golang/fermyon/spin/inbound-http"
)

// handler is the function that will be called by the http trigger in Spin.
var handler = defaultHandler

// defaultHandler is a placeholder for returning a useful error to stderr when
// the handler is not set.
var defaultHandler = func(http.ResponseWriter, *http.Request) {
	fmt.Fprintln(os.Stderr, "http handler undefined")
}

// Handle sets the handler function for the http trigger.
// It must be set in an init() function.
func Handle(fn func(http.ResponseWriter, *http.Request)) {
	resp := inboundhttp.HandleRequest(req)
	handler = fn
}
