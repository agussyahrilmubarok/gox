package exception

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpErrors(t *testing.T) {
	errDummy := errors.New("underlying error")

	tests := []struct {
		name     string
		createFn func(string, error) *Http
		code     int
		msg      string
	}{
		{"BadRequest", NewHTTPBadRequest, http.StatusBadRequest, "bad request"},
		{"Unauthorized", NewHTTPUnauthorized, http.StatusUnauthorized, "unauthorized"},
		{"Forbidden", NewHTTPForbidden, http.StatusForbidden, "forbidden"},
		{"NotFound", NewHTTPNotFound, http.StatusNotFound, "not found"},
		{"MethodNotAllowed", NewHTTPMethodNotAllowed, http.StatusMethodNotAllowed, "method not allowed"},
		{"Conflict", NewHTTPConflict, http.StatusConflict, "conflict"},
		{"UnprocessableEntity", NewHTTPUnprocessableEntity, http.StatusUnprocessableEntity, "unprocessable entity"},
		{"TooManyRequests", NewHTTPTooManyRequests, http.StatusTooManyRequests, "too many requests"},
		{"RequestTimeout", NewHTTPRequestTimeout, http.StatusRequestTimeout, "request timeout"},
		{"Internal", NewHTTPInternal, http.StatusInternalServerError, "internal server error"},
		{"BadGateway", NewHTTPBadGateway, http.StatusBadGateway, "bad gateway"},
		{"ServiceUnavailable", NewHTTPServiceUnavailable, http.StatusServiceUnavailable, "service unavailable"},
		{"GatewayTimeout", NewHTTPGatewayTimeout, http.StatusGatewayTimeout, "gateway timeout"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.createFn(tt.msg, errDummy)
			assert.NotNil(t, e)
			assert.Equal(t, tt.code, e.Code)
			assert.Equal(t, tt.msg, e.Message)
			assert.Equal(t, errDummy, e.Err)
			assert.Equal(t, tt.msg, e.Error())
		})
	}
}
