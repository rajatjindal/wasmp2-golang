package inboundhttp

type Impl struct{}

// Handle sets the handler function for the http trigger.
// It must be set in an init() function.
//
//export fermyon:spin/inbound-http@2.0.0#handle-request
func (i *Impl) HandleRequest(req Request) Response {
	return Response{
		Status: 401,
	}
}

func Export(impl Interface) {
	instance = impl
}
