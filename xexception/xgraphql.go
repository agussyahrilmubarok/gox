package xexception

// GraphQLError represents a structured GraphQL error with optional extensions.
type GraphQLError struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// NewGraphQLError creates a new GraphQLError with a message, code, and underlying error.
func NewGraphQLError(msg string, code string, err error) *GraphQLError {
	return &GraphQLError{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": code,
			"err":  err,
		},
	}
}

func NewGraphQLBadRequest(msg string, err error) *GraphQLError {
	return NewGraphQLError(msg, "BAD_REQUEST", err)
}

func NewGraphQLUnauthorized(msg string, err error) *GraphQLError {
	return NewGraphQLError(msg, "UNAUTHORIZED", err)
}

func NewGraphQLForbidden(msg string, err error) *GraphQLError {
	return NewGraphQLError(msg, "FORBIDDEN", err)
}

func NewGraphQLNotFound(msg string, err error) *GraphQLError {
	return NewGraphQLError(msg, "NOT_FOUND", err)
}

func NewGraphQLConflict(msg string, err error) *GraphQLError {
	return NewGraphQLError(msg, "CONFLICT", err)
}

func NewGraphQLUnprocessableEntity(msg string, err error) *GraphQLError {
	return NewGraphQLError(msg, "UNPROCESSABLE_ENTITY", err)
}

func NewGraphQLInternal(msg string, err error) *GraphQLError {
	return NewGraphQLError(msg, "INTERNAL_SERVER_ERROR", err)
}

func NewGraphQLServiceUnavailable(msg string, err error) *GraphQLError {
	return NewGraphQLError(msg, "SERVICE_UNAVAILABLE", err)
}
