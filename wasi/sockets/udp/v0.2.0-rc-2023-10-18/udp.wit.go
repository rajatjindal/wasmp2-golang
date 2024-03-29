// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package udp represents the interface "wasi:sockets/udp@0.2.0-rc-2023-10-18".
package udp

import (
	poll "github.com/rajatjindal/wasip2-golang/wasi/io/poll/v0.2.0-rc-2023-10-18"
	network "github.com/rajatjindal/wasip2-golang/wasi/sockets/network/v0.2.0-rc-2023-10-18"
	"github.com/ydnar/wasm-tools-go/cm"
)

// Datagram represents the record "wasi:sockets/udp@0.2.0-rc-2023-10-18#datagram".
//
//	record datagram {
//		data: list<u8>,
//		remote-address: ip-socket-address,
//	}
type Datagram struct {
	Data cm.List[uint8]

	// Theoretical max size: ~64 KiB. In practice, typically less than 1500 bytes.
	RemoteAddress IPSocketAddress
}

// ErrorCode represents the enum "wasi:sockets/network@0.2.0-rc-2023-10-18#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IPAddressFamily represents the enum "wasi:sockets/network@0.2.0-rc-2023-10-18#ip-address-family".
//
// See [network.IPAddressFamily] for more information.
type IPAddressFamily = network.IPAddressFamily

// IPSocketAddress represents the variant "wasi:sockets/network@0.2.0-rc-2023-10-18#ip-socket-address".
//
// See [network.IPSocketAddress] for more information.
type IPSocketAddress = network.IPSocketAddress

// Network represents the resource "wasi:sockets/network@0.2.0-rc-2023-10-18#network".
//
// See [network.Network] for more information.
type Network = network.Network

// Pollable represents the resource "wasi:io/poll@0.2.0-rc-2023-10-18#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// UDPSocket represents the resource "wasi:sockets/udp@0.2.0-rc-2023-10-18#udp-socket".
//
// A UDP socket handle.
//
//	resource udp-socket
type UDPSocket cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self UDPSocket) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [resource-drop]udp-socket
//go:noescape
func (self UDPSocket) wasmimport_ResourceDrop()

// AddressFamily represents method "address-family".
//
// Whether this is a IPv4 or IPv6 socket.
//
// Equivalent to the SO_DOMAIN socket option.
//
//	address-family: func() -> ip-address-family
//
//go:nosplit
func (self UDPSocket) AddressFamily() IPAddressFamily {
	return self.wasmimport_AddressFamily()
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.address-family
//go:noescape
func (self UDPSocket) wasmimport_AddressFamily() IPAddressFamily

// FinishBind represents method "finish-bind".
//
//	finish-bind: func() -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) FinishBind() cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_FinishBind(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.finish-bind
//go:noescape
func (self UDPSocket) wasmimport_FinishBind(result *cm.ErrResult[struct{}, ErrorCode])

// FinishConnect represents method "finish-connect".
//
//	finish-connect: func() -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) FinishConnect() cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_FinishConnect(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.finish-connect
//go:noescape
func (self UDPSocket) wasmimport_FinishConnect(result *cm.ErrResult[struct{}, ErrorCode])

// IPv6Only represents method "ipv6-only".
//
// Whether IPv4 compatibility (dual-stack) mode is disabled or not.
//
// Equivalent to the IPV6_V6ONLY socket option.
//
// # Typical errors
// - `not-supported`:        (get/set) `this` socket is an IPv4 socket.
// - `invalid-state`:        (set) The socket is already bound.
// - `not-supported`:        (set) Host does not support dual-stack sockets. (Implementations
// are not required to.)
//
//	ipv6-only: func() -> result<bool, error-code>
//
//go:nosplit
func (self UDPSocket) IPv6Only() cm.OKResult[bool, ErrorCode] {
	var result cm.OKResult[bool, ErrorCode]
	self.wasmimport_IPv6Only(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.ipv6-only
//go:noescape
func (self UDPSocket) wasmimport_IPv6Only(result *cm.OKResult[bool, ErrorCode])

// LocalAddress represents method "local-address".
//
// Get the current bound address.
//
// POSIX mentions:
// > If the socket has not been bound to a local name, the value
// > stored in the object pointed to by `address` is unspecified.
//
// WASI is stricter and requires `local-address` to return `invalid-state` when the
// socket hasn't been bound yet.
//
// # Typical errors
// - `invalid-state`: The socket is not bound to any local address.
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/getsockname.html>
// - <https://man7.org/linux/man-pages/man2/getsockname.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-getsockname>
// - <https://man.freebsd.org/cgi/man.cgi?getsockname>
//
//	local-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self UDPSocket) LocalAddress() cm.OKResult[IPSocketAddress, ErrorCode] {
	var result cm.OKResult[IPSocketAddress, ErrorCode]
	self.wasmimport_LocalAddress(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.local-address
//go:noescape
func (self UDPSocket) wasmimport_LocalAddress(result *cm.OKResult[IPSocketAddress, ErrorCode])

// Receive represents method "receive".
//
// Receive messages on the socket.
//
// This function attempts to receive up to `max-results` datagrams on the socket without
// blocking.
// The returned list may contain fewer elements than requested, but never more.
// If `max-results` is 0, this function returns successfully with an empty list.
//
// # Typical errors
// - `invalid-state`:      The socket is not bound to any local address. (EINVAL)
// - `remote-unreachable`: The remote address is not reachable. (ECONNREFUSED, ECONNRESET,
// ENETRESET on Windows, EHOSTUNREACH, EHOSTDOWN, ENETUNREACH, ENETDOWN)
// - `would-block`:        There is no pending data available to be read at the moment.
// (EWOULDBLOCK, EAGAIN)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/recvfrom.html>
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/recvmsg.html>
// - <https://man7.org/linux/man-pages/man2/recv.2.html>
// - <https://man7.org/linux/man-pages/man2/recvmmsg.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-recv>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-recvfrom>
// - <https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms741687(v=vs.85)>
// - <https://man.freebsd.org/cgi/man.cgi?query=recv&sektion=2>
//
//	receive: func(max-results: u64) -> result<list<datagram>, error-code>
//
//go:nosplit
func (self UDPSocket) Receive(maxResults uint64) cm.OKResult[cm.List[Datagram], ErrorCode] {
	var result cm.OKResult[cm.List[Datagram], ErrorCode]
	self.wasmimport_Receive(maxResults, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.receive
//go:noescape
func (self UDPSocket) wasmimport_Receive(maxResults uint64, result *cm.OKResult[cm.List[Datagram], ErrorCode])

// ReceiveBufferSize represents method "receive-buffer-size".
//
// The kernel buffer space reserved for sends/receives on this socket.
//
// Note #1: an implementation may choose to cap or round the buffer size when setting
// the value.
// In other words, after setting a value, reading the same setting back may return
// a different value.
//
// Note #2: there is not necessarily a direct relationship between the kernel buffer
// size and the bytes of
// actual data to be sent/received by the application, because the kernel might also
// use the buffer space
// for internal metadata structures.
//
// Equivalent to the SO_RCVBUF and SO_SNDBUF socket options.
//
//	receive-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self UDPSocket) ReceiveBufferSize() cm.OKResult[uint64, ErrorCode] {
	var result cm.OKResult[uint64, ErrorCode]
	self.wasmimport_ReceiveBufferSize(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.receive-buffer-size
//go:noescape
func (self UDPSocket) wasmimport_ReceiveBufferSize(result *cm.OKResult[uint64, ErrorCode])

// RemoteAddress represents method "remote-address".
//
// Get the address set with `connect`.
//
// # Typical errors
// - `invalid-state`: The socket is not connected to a remote address. (ENOTCONN)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/getpeername.html>
// - <https://man7.org/linux/man-pages/man2/getpeername.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-getpeername>
// - <https://man.freebsd.org/cgi/man.cgi?query=getpeername&sektion=2&n=1>
//
//	remote-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self UDPSocket) RemoteAddress() cm.OKResult[IPSocketAddress, ErrorCode] {
	var result cm.OKResult[IPSocketAddress, ErrorCode]
	self.wasmimport_RemoteAddress(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.remote-address
//go:noescape
func (self UDPSocket) wasmimport_RemoteAddress(result *cm.OKResult[IPSocketAddress, ErrorCode])

// Send represents method "send".
//
// Send messages on the socket.
//
// This function attempts to send all provided `datagrams` on the socket without blocking
// and
// returns how many messages were actually sent (or queued for sending).
//
// This function semantically behaves the same as iterating the `datagrams` list and
// sequentially
// sending each individual datagram until either the end of the list has been reached
// or the first error occurred.
// If at least one datagram has been sent successfully, this function never returns
// an error.
//
// If the input list is empty, the function returns `ok(0)`.
//
// The remote address option is required. To send a message to the "connected" peer,
// call `remote-address` to get their address.
//
// # Typical errors
// - `invalid-argument`:        The `remote-address` has the wrong address family.
// (EAFNOSUPPORT)
// - `invalid-argument`:        `remote-address` is a non-IPv4-mapped IPv6 address,
// but the socket was bound to a specific IPv4-mapped IPv6 address. (or vice versa)
// - `invalid-argument`:        The IP address in `remote-address` is set to INADDR_ANY
// (`0.0.0.0` / `::`). (EDESTADDRREQ, EADDRNOTAVAIL)
// - `invalid-argument`:        The port in `remote-address` is set to 0. (EDESTADDRREQ,
// EADDRNOTAVAIL)
// - `invalid-argument`:        The socket is in "connected" mode and the `datagram.remote-address`
// does not match the address passed to `connect`. (EISCONN)
// - `invalid-state`:           The socket is not bound to any local address. Unlike
// POSIX, this function does not perform an implicit bind.
// - `remote-unreachable`:      The remote address is not reachable. (ECONNREFUSED,
// ECONNRESET, ENETRESET on Windows, EHOSTUNREACH, EHOSTDOWN, ENETUNREACH, ENETDOWN)
// - `datagram-too-large`:      The datagram is too large. (EMSGSIZE)
// - `would-block`:             The send buffer is currently full. (EWOULDBLOCK, EAGAIN)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/sendto.html>
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/sendmsg.html>
// - <https://man7.org/linux/man-pages/man2/send.2.html>
// - <https://man7.org/linux/man-pages/man2/sendmmsg.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-send>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-sendto>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-wsasendmsg>
// - <https://man.freebsd.org/cgi/man.cgi?query=send&sektion=2>
//
//	send: func(datagrams: list<datagram>) -> result<u64, error-code>
//
//go:nosplit
func (self UDPSocket) Send(datagrams cm.List[Datagram]) cm.OKResult[uint64, ErrorCode] {
	var result cm.OKResult[uint64, ErrorCode]
	self.wasmimport_Send(datagrams, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.send
//go:noescape
func (self UDPSocket) wasmimport_Send(datagrams cm.List[Datagram], result *cm.OKResult[uint64, ErrorCode])

// SendBufferSize represents method "send-buffer-size".
//
//	send-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self UDPSocket) SendBufferSize() cm.OKResult[uint64, ErrorCode] {
	var result cm.OKResult[uint64, ErrorCode]
	self.wasmimport_SendBufferSize(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.send-buffer-size
//go:noescape
func (self UDPSocket) wasmimport_SendBufferSize(result *cm.OKResult[uint64, ErrorCode])

// SetIPv6Only represents method "set-ipv6-only".
//
//	set-ipv6-only: func(value: bool) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetIPv6Only(value bool) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_SetIPv6Only(value, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.set-ipv6-only
//go:noescape
func (self UDPSocket) wasmimport_SetIPv6Only(value bool, result *cm.ErrResult[struct{}, ErrorCode])

// SetReceiveBufferSize represents method "set-receive-buffer-size".
//
//	set-receive-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetReceiveBufferSize(value uint64) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_SetReceiveBufferSize(value, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.set-receive-buffer-size
//go:noescape
func (self UDPSocket) wasmimport_SetReceiveBufferSize(value uint64, result *cm.ErrResult[struct{}, ErrorCode])

// SetSendBufferSize represents method "set-send-buffer-size".
//
//	set-send-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetSendBufferSize(value uint64) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_SetSendBufferSize(value, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.set-send-buffer-size
//go:noescape
func (self UDPSocket) wasmimport_SetSendBufferSize(value uint64, result *cm.ErrResult[struct{}, ErrorCode])

// SetUnicastHopLimit represents method "set-unicast-hop-limit".
//
//	set-unicast-hop-limit: func(value: u8) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetUnicastHopLimit(value uint8) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_SetUnicastHopLimit(value, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.set-unicast-hop-limit
//go:noescape
func (self UDPSocket) wasmimport_SetUnicastHopLimit(value uint8, result *cm.ErrResult[struct{}, ErrorCode])

// StartBind represents method "start-bind".
//
// Bind the socket to a specific network on the provided IP address and port.
//
// If the IP address is zero (`0.0.0.0` in IPv4, `::` in IPv6), it is left to the
// implementation to decide which
// network interface(s) to bind to.
// If the TCP/UDP port is zero, the socket will be bound to a random free port.
//
// When a socket is not explicitly bound, the first invocation to connect will implicitly
// bind the socket.
//
// Unlike in POSIX, this function is async. This enables interactive WASI hosts to
// inject permission prompts.
//
// # Typical `start` errors
// - `invalid-argument`:          The `local-address` has the wrong address family.
// (EAFNOSUPPORT, EFAULT on Windows)
// - `invalid-state`:             The socket is already bound. (EINVAL)
//
// # Typical `finish` errors
// - `address-in-use`:            No ephemeral ports available. (EADDRINUSE, ENOBUFS
// on Windows)
// - `address-in-use`:            Address is already in use. (EADDRINUSE)
// - `address-not-bindable`:      `local-address` is not an address that the `network`
// can bind to. (EADDRNOTAVAIL)
// - `not-in-progress`:           A `bind` operation is not in progress.
// - `would-block`:               Can't finish the operation, it is still in progress.
// (EWOULDBLOCK, EAGAIN)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/bind.html>
// - <https://man7.org/linux/man-pages/man2/bind.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-bind>
// - <https://man.freebsd.org/cgi/man.cgi?query=bind&sektion=2&format=html>
//
//	start-bind: func(network: borrow<network>, local-address: ip-socket-address) ->
//	result<_, error-code>
//
//go:nosplit
func (self UDPSocket) StartBind(network_ Network, localAddress IPSocketAddress) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_StartBind(network_, localAddress, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.start-bind
//go:noescape
func (self UDPSocket) wasmimport_StartBind(network_ Network, localAddress IPSocketAddress, result *cm.ErrResult[struct{}, ErrorCode])

// StartConnect represents method "start-connect".
//
// Set the destination address.
//
// The local-address is updated based on the best network path to `remote-address`.
//
// When a destination address is set:
// - all receive operations will only return datagrams sent from the provided `remote-address`.
// - the `send` function can only be used to send to this destination.
//
// Note that this function does not generate any network traffic and the peer is not
// aware of this "connection".
//
// Unlike in POSIX, this function is async. This enables interactive WASI hosts to
// inject permission prompts.
//
// # Typical `start` errors
// - `invalid-argument`:          The `remote-address` has the wrong address family.
// (EAFNOSUPPORT)
// - `invalid-argument`:          `remote-address` is a non-IPv4-mapped IPv6 address,
// but the socket was bound to a specific IPv4-mapped IPv6 address. (or vice versa)
// - `invalid-argument`:          The IP address in `remote-address` is set to INADDR_ANY
// (`0.0.0.0` / `::`). (EDESTADDRREQ, EADDRNOTAVAIL)
// - `invalid-argument`:          The port in `remote-address` is set to 0. (EDESTADDRREQ,
// EADDRNOTAVAIL)
// - `invalid-argument`:          The socket is already bound to a different network.
// The `network` passed to `connect` must be identical to the one passed to `bind`.
//
// # Typical `finish` errors
// - `address-in-use`:            Tried to perform an implicit bind, but there were
// no ephemeral ports available. (EADDRINUSE, EADDRNOTAVAIL on Linux, EAGAIN on BSD)
// - `not-in-progress`:           A `connect` operation is not in progress.
// - `would-block`:               Can't finish the operation, it is still in progress.
// (EWOULDBLOCK, EAGAIN)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/connect.html>
// - <https://man7.org/linux/man-pages/man2/connect.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-connect>
// - <https://man.freebsd.org/cgi/man.cgi?connect>
//
//	start-connect: func(network: borrow<network>, remote-address: ip-socket-address)
//	-> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) StartConnect(network_ Network, remoteAddress IPSocketAddress) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_StartConnect(network_, remoteAddress, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.start-connect
//go:noescape
func (self UDPSocket) wasmimport_StartConnect(network_ Network, remoteAddress IPSocketAddress, result *cm.ErrResult[struct{}, ErrorCode])

// Subscribe represents method "subscribe".
//
// Create a `pollable` which will resolve once the socket is ready for I/O.
//
// Note: this function is here for WASI Preview2 only.
// It's planned to be removed when `future` is natively supported in Preview3.
//
//	subscribe: func() -> own<pollable>
//
//go:nosplit
func (self UDPSocket) Subscribe() Pollable {
	return self.wasmimport_Subscribe()
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.subscribe
//go:noescape
func (self UDPSocket) wasmimport_Subscribe() Pollable

// UnicastHopLimit represents method "unicast-hop-limit".
//
// Equivalent to the IP_TTL & IPV6_UNICAST_HOPS socket options.
//
//	unicast-hop-limit: func() -> result<u8, error-code>
//
//go:nosplit
func (self UDPSocket) UnicastHopLimit() cm.OKResult[uint8, ErrorCode] {
	var result cm.OKResult[uint8, ErrorCode]
	self.wasmimport_UnicastHopLimit(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0-rc-2023-10-18 [method]udp-socket.unicast-hop-limit
//go:noescape
func (self UDPSocket) wasmimport_UnicastHopLimit(result *cm.OKResult[uint8, ErrorCode])
