package constants

import "net/textproto"

var (
	XServiceName = textproto.CanonicalMIMEHeaderKey("x-service-name")
	XApiKey      = textproto.CanonicalMIMEHeaderKey("x-api-key")
	XReqestAt    = textproto.CanonicalMIMEHeaderKey("x-request-at")
	Authrization = textproto.CanonicalMIMEHeaderKey("authorization")
)
