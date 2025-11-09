package exception

import "net/http"

// Http represents a structured HTTP error with status code, message, and optional wrapped error.
type Http struct {
	Code    int    // HTTP status code
	Message string // Human-readable error message
	Err     error  // Wrapped error, if any
}

// Error implements the error interface.
func (e *Http) Error() string {
	return e.Message
}

func NewHTTPBadRequest(msg string, err error) *Http {
	return &Http{Code: http.StatusBadRequest, Message: msg, Err: err}
}

func NewHTTPUnauthorized(msg string, err error) *Http {
	return &Http{Code: http.StatusUnauthorized, Message: msg, Err: err}
}

func NewHTTPForbidden(msg string, err error) *Http {
	return &Http{Code: http.StatusForbidden, Message: msg, Err: err}
}

func NewHTTPNotFound(msg string, err error) *Http {
	return &Http{Code: http.StatusNotFound, Message: msg, Err: err}
}

func NewHTTPMethodNotAllowed(msg string, err error) *Http {
	return &Http{Code: http.StatusMethodNotAllowed, Message: msg, Err: err}
}

func NewHTTPConflict(msg string, err error) *Http {
	return &Http{Code: http.StatusConflict, Message: msg, Err: err}
}

func NewHTTPUnprocessableEntity(msg string, err error) *Http {
	return &Http{Code: http.StatusUnprocessableEntity, Message: msg, Err: err}
}

func NewHTTPTooManyRequests(msg string, err error) *Http {
	return &Http{Code: http.StatusTooManyRequests, Message: msg, Err: err}
}

func NewHTTPRequestTimeout(msg string, err error) *Http {
	return &Http{Code: http.StatusRequestTimeout, Message: msg, Err: err}
}

func NewHTTPInternal(msg string, err error) *Http {
	return &Http{Code: http.StatusInternalServerError, Message: msg, Err: err}
}

func NewHTTPBadGateway(msg string, err error) *Http {
	return &Http{Code: http.StatusBadGateway, Message: msg, Err: err}
}

func NewHTTPServiceUnavailable(msg string, err error) *Http {
	return &Http{Code: http.StatusServiceUnavailable, Message: msg, Err: err}
}

func NewHTTPGatewayTimeout(msg string, err error) *Http {
	return &Http{Code: http.StatusGatewayTimeout, Message: msg, Err: err}
}

func NewHTTPPaymentRequired(msg string, err error) *Http {
	return &Http{Code: http.StatusPaymentRequired, Message: msg, Err: err}
}

func NewHTTPNotAcceptable(msg string, err error) *Http {
	return &Http{Code: http.StatusNotAcceptable, Message: msg, Err: err}
}

func NewHTTPProxyAuthRequired(msg string, err error) *Http {
	return &Http{Code: http.StatusProxyAuthRequired, Message: msg, Err: err}
}

func NewHTTPGone(msg string, err error) *Http {
	return &Http{Code: http.StatusGone, Message: msg, Err: err}
}

func NewHTTPLengthRequired(msg string, err error) *Http {
	return &Http{Code: http.StatusLengthRequired, Message: msg, Err: err}
}

func NewHTTPPreconditionFailed(msg string, err error) *Http {
	return &Http{Code: http.StatusPreconditionFailed, Message: msg, Err: err}
}

func NewHTTPRequestEntityTooLarge(msg string, err error) *Http {
	return &Http{Code: http.StatusRequestEntityTooLarge, Message: msg, Err: err}
}

func NewHTTPRequestURITooLong(msg string, err error) *Http {
	return &Http{Code: http.StatusRequestURITooLong, Message: msg, Err: err}
}

func NewHTTPUnsupportedMediaType(msg string, err error) *Http {
	return &Http{Code: http.StatusUnsupportedMediaType, Message: msg, Err: err}
}

func NewHTTPExpectationFailed(msg string, err error) *Http {
	return &Http{Code: http.StatusExpectationFailed, Message: msg, Err: err}
}

func NewHTTPLocked(msg string, err error) *Http {
	return &Http{Code: http.StatusLocked, Message: msg, Err: err}
}

func NewHTTPFailedDependency(msg string, err error) *Http {
	return &Http{Code: http.StatusFailedDependency, Message: msg, Err: err}
}

func NewHTTPNotImplemented(msg string, err error) *Http {
	return &Http{Code: http.StatusNotImplemented, Message: msg, Err: err}
}

func NewHTTPHTTPVersionNotSupported(msg string, err error) *Http {
	return &Http{Code: http.StatusHTTPVersionNotSupported, Message: msg, Err: err}
}

func NewHTTPVariantAlsoNegotiates(msg string, err error) *Http {
	return &Http{Code: http.StatusVariantAlsoNegotiates, Message: msg, Err: err}
}

func NewHTTPInsufficientStorage(msg string, err error) *Http {
	return &Http{Code: http.StatusInsufficientStorage, Message: msg, Err: err}
}

func NewHTTPLoopDetected(msg string, err error) *Http {
	return &Http{Code: http.StatusLoopDetected, Message: msg, Err: err}
}

func NewHTTPNotExtended(msg string, err error) *Http {
	return &Http{Code: http.StatusNotExtended, Message: msg, Err: err}
}

func NewHTTPNetworkAuthenticationRequired(msg string, err error) *Http {
	return &Http{Code: http.StatusNetworkAuthenticationRequired, Message: msg, Err: err}
}
