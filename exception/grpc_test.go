package exception

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPCErrors(t *testing.T) {
	errDummy := errors.New("underlying error")

	tests := []struct {
		name     string
		createFn func(string, error) error
		code     codes.Code
		msg      string
	}{
		{"NotFound", NewGRPCNotFound, codes.NotFound, "not found"},
		{"Internal", NewGRPCInternal, codes.Internal, "internal error"},
		{"Canceled", NewGRPCCanceled, codes.Canceled, "canceled"},
		{"Unknown", NewGRPCUnknown, codes.Unknown, "unknown error"},
		{"InvalidArgument", NewGRPCInvalidArgument, codes.InvalidArgument, "invalid argument"},
		{"DeadlineExceeded", NewGRPCDeadlineExceeded, codes.DeadlineExceeded, "deadline exceeded"},
		{"AlreadyExists", NewGRPCAlreadyExists, codes.AlreadyExists, "already exists"},
		{"PermissionDenied", NewGRPCPermissionDenied, codes.PermissionDenied, "permission denied"},
		{"ResourceExhausted", NewGRPCResourceExhausted, codes.ResourceExhausted, "resource exhausted"},
		{"FailedPrecondition", NewGRPCFailedPrecondition, codes.FailedPrecondition, "failed precondition"},
		{"Aborted", NewGRPCAborted, codes.Aborted, "aborted"},
		{"OutOfRange", NewGRPCOutOfRange, codes.OutOfRange, "out of range"},
		{"Unimplemented", NewGRPCUnimplemented, codes.Unimplemented, "unimplemented"},
		{"Unavailable", NewGRPCUnavailable, codes.Unavailable, "unavailable"},
		{"DataLoss", NewGRPCDataLoss, codes.DataLoss, "data loss"},
		{"Unauthenticated", NewGRPCUnauthenticated, codes.Unauthenticated, "unauthenticated"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.createFn(tt.msg, errDummy)
			assert.NotNil(t, e)

			st := status.Convert(e)
			assert.Equal(t, tt.code, st.Code())
			assert.Contains(t, st.Message(), tt.msg)
			assert.Contains(t, st.Message(), errDummy.Error())
		})
	}
}
