package xexception_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/agussyahrilmubarok/gohelp/xexception"
	"github.com/stretchr/testify/assert"
)

func TestHttpErrors(t *testing.T) {
	errDummy := errors.New("underlying error")

	tests := []struct {
		name     string
		createFn func(string, error) *xexception.Http
		code     int
		msg      string
	}{
		{"BadRequest", xexception.NewHTTPBadRequest, http.StatusBadRequest, "bad request"},
		{"Unauthorized", xexception.NewHTTPUnauthorized, http.StatusUnauthorized, "unauthorized"},
		{"Forbidden", xexception.NewHTTPForbidden, http.StatusForbidden, "forbidden"},
		{"NotFound", xexception.NewHTTPNotFound, http.StatusNotFound, "not found"},
		{"MethodNotAllowed", xexception.NewHTTPMethodNotAllowed, http.StatusMethodNotAllowed, "method not allowed"},
		{"Conflict", xexception.NewHTTPConflict, http.StatusConflict, "conflict"},
		{"UnprocessableEntity", xexception.NewHTTPUnprocessableEntity, http.StatusUnprocessableEntity, "unprocessable entity"},
		{"TooManyRequests", xexception.NewHTTPTooManyRequests, http.StatusTooManyRequests, "too many requests"},
		{"RequestTimeout", xexception.NewHTTPRequestTimeout, http.StatusRequestTimeout, "request timeout"},
		{"Internal", xexception.NewHTTPInternal, http.StatusInternalServerError, "internal server error"},
		{"BadGateway", xexception.NewHTTPBadGateway, http.StatusBadGateway, "bad gateway"},
		{"ServiceUnavailable", xexception.NewHTTPServiceUnavailable, http.StatusServiceUnavailable, "service unavailable"},
		{"GatewayTimeout", xexception.NewHTTPGatewayTimeout, http.StatusGatewayTimeout, "gateway timeout"},
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
