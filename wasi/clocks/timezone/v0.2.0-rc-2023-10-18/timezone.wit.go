// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package timezone represents the interface "wasi:clocks/timezone@0.2.0-rc-2023-10-18".
package timezone

import (
	wallclock "github.com/rajatjindal/wasip2-golang/wasi/clocks/wall-clock/v0.2.0-rc-2023-10-18"
)

// DateTime represents the record "wasi:clocks/wall-clock@0.2.0-rc-2023-10-18#datetime".
//
// See [wallclock.DateTime] for more information.
type DateTime = wallclock.DateTime

// TimezoneDisplay represents the record "wasi:clocks/timezone@0.2.0-rc-2023-10-18#timezone-display".
//
// Information useful for displaying the timezone of a specific `datetime`.
//
// This information may vary within a single `timezone` to reflect daylight
// saving time adjustments.
//
//	record timezone-display {
//		utc-offset: s32,
//		name: string,
//		in-daylight-saving-time: bool,
//	}
type TimezoneDisplay struct {
	// The number of seconds difference between UTC time and the local
	// time of the timezone.
	//
	// The returned value will always be less than 86400 which is the
	// number of seconds in a day (24*60*60).
	//
	// In implementations that do not expose an actual time zone, this
	// should return 0.
	UtcOffset int32

	// The abbreviated name of the timezone to display to a user. The name
	// `UTC` indicates Coordinated Universal Time. Otherwise, this should
	// reference local standards for the name of the time zone.
	//
	// In implementations that do not expose an actual time zone, this
	// should be the string `UTC`.
	//
	// In time zones that do not have an applicable name, a formatted
	// representation of the UTC offset may be returned, such as `-04:00`.
	Name string

	// Whether daylight saving time is active.
	//
	// In implementations that do not expose an actual time zone, this
	// should return false.
	InDaylightSavingTime bool
}

// Display represents function "wasi:clocks/timezone@0.2.0-rc-2023-10-18#display".
//
// Return information needed to display the given `datetime`. This includes
// the UTC offset, the time zone name, and a flag indicating whether
// daylight saving time is active.
//
// If the timezone cannot be determined for the given `datetime`, return a
// `timezone-display` for `UTC` with a `utc-offset` of 0 and no daylight
// saving time.
//
//	display: func(when: datetime) -> timezone-display
//
//go:nosplit
func Display(when DateTime) TimezoneDisplay {
	var result TimezoneDisplay
	wasmimport_Display(when, &result)
	return result
}

//go:wasmimport wasi:clocks/timezone@0.2.0-rc-2023-10-18 display
//go:noescape
func wasmimport_Display(when DateTime, result *TimezoneDisplay)

// UtcOffset represents function "wasi:clocks/timezone@0.2.0-rc-2023-10-18#utc-offset".
//
// The same as `display`, but only return the UTC offset.
//
//	utc-offset: func(when: datetime) -> s32
//
//go:nosplit
func UtcOffset(when DateTime) int32 {
	return wasmimport_UtcOffset(when)
}

//go:wasmimport wasi:clocks/timezone@0.2.0-rc-2023-10-18 utc-offset
//go:noescape
func wasmimport_UtcOffset(when DateTime) int32
