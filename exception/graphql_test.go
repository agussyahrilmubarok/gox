package exception

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphQLErrors(t *testing.T) {
	errDummy := errors.New("underlying error")

	tests := []struct {
		name     string
		createFn func(string, error) *GraphQLError
		code     string
		msg      string
	}{
		{"BadRequest", NewGraphQLBadRequest, "BAD_REQUEST", "bad request"},
		{"Unauthorized", NewGraphQLUnauthorized, "UNAUTHORIZED", "unauthorized"},
		{"Forbidden", NewGraphQLForbidden, "FORBIDDEN", "forbidden"},
		{"NotFound", NewGraphQLNotFound, "NOT_FOUND", "not found"},
		{"Conflict", NewGraphQLConflict, "CONFLICT", "conflict"},
		{"UnprocessableEntity", NewGraphQLUnprocessableEntity, "UNPROCESSABLE_ENTITY", "unprocessable entity"},
		{"Internal", NewGraphQLInternal, "INTERNAL_SERVER_ERROR", "internal server error"},
		{"ServiceUnavailable", NewGraphQLServiceUnavailable, "SERVICE_UNAVAILABLE", "service unavailable"},
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
