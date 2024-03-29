// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package monotonicclock represents the interface "wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10".
//
// WASI Monotonic Clock is a clock API intended to let users measure elapsed
// time.
//
// It is intended to be portable at least between Unix-family platforms and
// Windows.
//
// A monotonic clock is a clock which has an unspecified initial value, and
// successive reads of the clock will produce non-decreasing values.
//
// It is intended for measuring elapsed time.
package monotonicclock

import (
	poll "github.com/rajatjindal/wasip2-golang/wasi/io/poll/v0.2.0-rc-2023-11-10"
)

// Duration represents the type "wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10#duration".
//
// A duration of time, in nanoseconds.
//
//	type duration = u64
type Duration uint64

// Instant represents the type "wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10#instant".
//
// An instant in time, in nanoseconds. An instant is relative to an
// unspecified initial value, and can only be compared to instances from
// the same monotonic-clock.
//
//	type instant = u64
type Instant uint64

// Pollable represents the resource "wasi:io/poll@0.2.0-rc-2023-11-10#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// Now represents function "wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10#now".
//
// Read the current value of the clock.
//
// The clock is monotonic, therefore calling this function repeatedly will
// produce a sequence of non-decreasing values.
//
//	now: func() -> instant
//
//go:nosplit
func Now() Instant {
	return wasmimport_Now()
}

//go:wasmimport wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10 now
//go:noescape
func wasmimport_Now() Instant

// Resolution represents function "wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10#resolution".
//
// Query the resolution of the clock. Returns the duration of time
// corresponding to a clock tick.
//
//	resolution: func() -> duration
//
//go:nosplit
func Resolution() Duration {
	return wasmimport_Resolution()
}

//go:wasmimport wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10 resolution
//go:noescape
func wasmimport_Resolution() Duration

// SubscribeDuration represents function "wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10#subscribe-duration".
//
// Create a `pollable` which will resolve once the given duration has
// elapsed, starting at the time at which this function was called.
// occured.
//
//	subscribe-duration: func(when: duration) -> own<pollable>
//
//go:nosplit
func SubscribeDuration(when Duration) Pollable {
	return wasmimport_SubscribeDuration(when)
}

//go:wasmimport wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10 subscribe-duration
//go:noescape
func wasmimport_SubscribeDuration(when Duration) Pollable

// SubscribeInstant represents function "wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10#subscribe-instant".
//
// Create a `pollable` which will resolve once the specified instant
// occured.
//
//	subscribe-instant: func(when: instant) -> own<pollable>
//
//go:nosplit
func SubscribeInstant(when Instant) Pollable {
	return wasmimport_SubscribeInstant(when)
}

//go:wasmimport wasi:clocks/monotonic-clock@0.2.0-rc-2023-11-10 subscribe-instant
//go:noescape
func wasmimport_SubscribeInstant(when Instant) Pollable
