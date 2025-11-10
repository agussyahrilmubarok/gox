package xexception_test

import (
	"errors"
	"testing"

	"github.com/agussyahrilmubarok/gohelp/xexception"
	"github.com/stretchr/testify/assert"
)

func TestGraphQLErrors(t *testing.T) {
	errDummy := errors.New("underlying error")

	tests := []struct {
		name     string
		createFn func(string, error) *xexception.GraphQLError
		code     string
		msg      string
	}{
		{"BadRequest", xexception.NewGraphQLBadRequest, "BAD_REQUEST", "bad request"},
		{"Unauthorized", xexception.NewGraphQLUnauthorized, "UNAUTHORIZED", "unauthorized"},
		{"Forbidden", xexception.NewGraphQLForbidden, "FORBIDDEN", "forbidden"},
		{"NotFound", xexception.NewGraphQLNotFound, "NOT_FOUND", "not found"},
		{"Conflict", xexception.NewGraphQLConflict, "CONFLICT", "conflict"},
		{"UnprocessableEntity", xexception.NewGraphQLUnprocessableEntity, "UNPROCESSABLE_ENTITY", "unprocessable entity"},
		{"Internal", xexception.NewGraphQLInternal, "INTERNAL_SERVER_ERROR", "internal server error"},
		{"ServiceUnavailable", xexception.NewGraphQLServiceUnavailable, "SERVICE_UNAVAILABLE", "service unavailable"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.createFn(tt.msg, errDummy)
			assert.NotNil(t, e)
			assert.Equal(t, tt.msg, e.Message)
			assert.Equal(t, tt.code, e.Extensions["code"])
			assert.Equal(t, errDummy, e.Extensions["err"])
		})
	}
}
