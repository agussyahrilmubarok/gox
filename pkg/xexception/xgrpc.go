package xexception

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGRPCNotFound(msg string, err error) error {
	return status.Errorf(codes.NotFound, "%s: %v", msg, err)
}

func NewGRPCInternal(msg string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", msg, err)
}

func NewGRPCCanceled(msg string, err error) error {
	return status.Errorf(codes.Canceled, "%s: %v", msg, err)
}

func NewGRPCUnknown(msg string, err error) error {
	return status.Errorf(codes.Unknown, "%s: %v", msg, err)
}

func NewGRPCInvalidArgument(msg string, err error) error {
	return status.Errorf(codes.InvalidArgument, "%s: %v", msg, err)
}

func NewGRPCDeadlineExceeded(msg string, err error) error {
	return status.Errorf(codes.DeadlineExceeded, "%s: %v", msg, err)
}

func NewGRPCAlreadyExists(msg string, err error) error {
	return status.Errorf(codes.AlreadyExists, "%s: %v", msg, err)
}

func NewGRPCPermissionDenied(msg string, err error) error {
	return status.Errorf(codes.PermissionDenied, "%s: %v", msg, err)
}

func NewGRPCResourceExhausted(msg string, err error) error {
	return status.Errorf(codes.ResourceExhausted, "%s: %v", msg, err)
}

func NewGRPCFailedPrecondition(msg string, err error) error {
	return status.Errorf(codes.FailedPrecondition, "%s: %v", msg, err)
}

func NewGRPCAborted(msg string, err error) error {
	return status.Errorf(codes.Aborted, "%s: %v", msg, err)
}

func NewGRPCOutOfRange(msg string, err error) error {
	return status.Errorf(codes.OutOfRange, "%s: %v", msg, err)
}

func NewGRPCUnimplemented(msg string, err error) error {
	return status.Errorf(codes.Unimplemented, "%s: %v", msg, err)
}

func NewGRPCUnavailable(msg string, err error) error {
	return status.Errorf(codes.Unavailable, "%s: %v", msg, err)
}

func NewGRPCDataLoss(msg string, err error) error {
	return status.Errorf(codes.DataLoss, "%s: %v", msg, err)
}

func NewGRPCUnauthenticated(msg string, err error) error {
	return status.Errorf(codes.Unauthenticated, "%s: %v", msg, err)
}
