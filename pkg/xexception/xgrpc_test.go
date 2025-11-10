package xexception_test

import (
	"errors"
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xexception"
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
		{"NotFound", xexception.NewGRPCNotFound, codes.NotFound, "not found"},
		{"Internal", xexception.NewGRPCInternal, codes.Internal, "internal error"},
		{"Canceled", xexception.NewGRPCCanceled, codes.Canceled, "canceled"},
		{"Unknown", xexception.NewGRPCUnknown, codes.Unknown, "unknown error"},
		{"InvalidArgument", xexception.NewGRPCInvalidArgument, codes.InvalidArgument, "invalid argument"},
		{"DeadlineExceeded", xexception.NewGRPCDeadlineExceeded, codes.DeadlineExceeded, "deadline exceeded"},
		{"AlreadyExists", xexception.NewGRPCAlreadyExists, codes.AlreadyExists, "already exists"},
		{"PermissionDenied", xexception.NewGRPCPermissionDenied, codes.PermissionDenied, "permission denied"},
		{"ResourceExhausted", xexception.NewGRPCResourceExhausted, codes.ResourceExhausted, "resource exhausted"},
		{"FailedPrecondition", xexception.NewGRPCFailedPrecondition, codes.FailedPrecondition, "failed precondition"},
		{"Aborted", xexception.NewGRPCAborted, codes.Aborted, "aborted"},
		{"OutOfRange", xexception.NewGRPCOutOfRange, codes.OutOfRange, "out of range"},
		{"Unimplemented", xexception.NewGRPCUnimplemented, codes.Unimplemented, "unimplemented"},
		{"Unavailable", xexception.NewGRPCUnavailable, codes.Unavailable, "unavailable"},
		{"DataLoss", xexception.NewGRPCDataLoss, codes.DataLoss, "data loss"},
		{"Unauthenticated", xexception.NewGRPCUnauthenticated, codes.Unauthenticated, "unauthenticated"},
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
