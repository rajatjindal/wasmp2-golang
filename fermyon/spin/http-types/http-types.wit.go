// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package httptypes represents the interface "fermyon:spin/http-types@2.0.0".
package httptypes

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Body represents the list "fermyon:spin/http-types@2.0.0#body".
//
//	type body = list<u8>
type Body cm.List[uint8]

// Headers represents the list "fermyon:spin/http-types@2.0.0#headers".
//
//	type headers = list<tuple<string, string>>
type Headers cm.List[[2]string]

// HTTPError represents the enum "fermyon:spin/http-types@2.0.0#http-error".
//
//	enum http-error {
//		success,
//		destination-not-allowed,
//		invalid-url,
//		request-error,
//		runtime-error,
//		too-many-requests
//	}
type HTTPError uint8

const (
	HTTPErrorSuccess HTTPError = iota
	HTTPErrorDestinationNotAllowed
	HTTPErrorInvalidURL
	HTTPErrorRequestError
	HTTPErrorRuntimeError
	HTTPErrorTooManyRequests
)

// HTTPStatus represents the type "fermyon:spin/http-types@2.0.0#http-status".
//
//	type http-status = u16
type HTTPStatus uint16

// Method represents the enum "fermyon:spin/http-types@2.0.0#method".
//
//	enum method {
//		get,
//		post,
//		put,
//		delete,
//		patch,
//		head,
//		options
//	}
type Method uint8

const (
	MethodGet Method = iota
	MethodPost
	MethodPut
	MethodDelete
	MethodPatch
	MethodHead
	MethodOptions
)

// Params represents the list "fermyon:spin/http-types@2.0.0#params".
//
//	type params = list<tuple<string, string>>
type Params cm.List[[2]string]

// Request represents the record "fermyon:spin/http-types@2.0.0#request".
//
//	record request {
//		method: method,
//		uri: uri,
//		headers: headers,
//		params: params,
//		body: option<body>,
//	}
type Request struct {
	Method  Method
	URI     URI
	Headers Headers
	Params  Params
	Body    cm.Option[Body]
}

// Response represents the record "fermyon:spin/http-types@2.0.0#response".
//
//	record response {
//		status: http-status,
//		headers: option<headers>,
//		body: option<body>,
//	}
type Response struct {
	Status  HTTPStatus
	Headers cm.Option[Headers]
	Body    cm.Option[Body]
}

// URI represents the type "fermyon:spin/http-types@2.0.0#uri".
//
//	type uri = string
type URI string
