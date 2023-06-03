package common

import "errors"

var (
	GenericHTTPServerErr = errors.New("the server has failed to process the request, please do not retry")
	ErrAccessDenied      = errors.New("access denied")
)
