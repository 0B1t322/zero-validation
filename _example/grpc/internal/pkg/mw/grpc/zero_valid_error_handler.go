package grpcmw

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	verrors "github.com/0B1t322/zero-validation/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"

	"google.golang.org/grpc"
)

type zeroValidationErrorHandlerTyped interface {
	Handle(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	isZeroValidationErrorHandler()
}

func ZeroValidationErrorHandler(typed zeroValidationErrorHandlerTyped) grpc.UnaryServerInterceptor {
	return typed.Handle
}

type zeroValidationErrorHandlerWithJsonFormat struct{}

func NewZeroValidationErrorHandlerWithJsonFormat() zeroValidationErrorHandlerTyped {
	return &zeroValidationErrorHandlerWithJsonFormat{}
}

func (z *zeroValidationErrorHandlerWithJsonFormat) Handle(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil && errors.Is(err, verrors.Errors{}) {
		buff := &bytes.Buffer{}
		errEncode := json.NewEncoder(buff).Encode(err)
		if errEncode != nil {
			return nil, status.Error(codes.InvalidArgument, "невалидные входные данные")
		}

		return nil, status.Error(codes.InvalidArgument, buff.String())
	}

	return resp, err
}

func (*zeroValidationErrorHandlerWithJsonFormat) isZeroValidationErrorHandler() {}

type zeroValidationErrorHandlerWithOneError struct{}

func NewZeroValidationErrorHandlerWithOneError() zeroValidationErrorHandlerTyped {
	return &zeroValidationErrorHandlerWithOneError{}
}

func (z *zeroValidationErrorHandlerWithOneError) Handle(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil && errors.Is(err, verrors.Errors{}) {
		return nil, status.Error(codes.InvalidArgument, getOneError(err))
	}

	return resp, err
}

func getOneError(err error) string {
	var vErrors verrors.Errors

	errors.As(err, &vErrors)

	return getOneErrorFromVerror(vErrors)
}

func getOneErrorFromVerror(vError verrors.Errors) string {
	for field, err := range vError {
		switch v := err.(type) {
		case verrors.Errors:
			return getOneErrorFromVerror(v)
		default:
			return field + ": " + err.Error()
		}
	}

	return ""
}

func (*zeroValidationErrorHandlerWithOneError) isZeroValidationErrorHandler() {}

type zeroValidationPathErrorHandler struct{}

func NewZeroValidationPathErrorHandler() zeroValidationErrorHandlerTyped {
	return &zeroValidationPathErrorHandler{}
}

func (z zeroValidationPathErrorHandler) Handle(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil && errors.Is(err, verrors.Errors{}) {
		verr := verrors.Errors{}
		errors.As(err, &verr)
		sb := strings.Builder{}
		buildPathError(verr, &sb)

		return nil, status.Error(codes.InvalidArgument, sb.String())
	}

	return resp, err
}

func (z zeroValidationPathErrorHandler) isZeroValidationErrorHandler() {}

func buildPathError(errors verrors.Errors, sb *strings.Builder) {
	for field, verr := range errors {
		sb.WriteString(field)

		switch v := verr.(type) {
		case verrors.ErrorSlice:
			if verr, isVerr := v[0].(verrors.Errors); isVerr {
				sb.WriteString(".")
				buildPathError(verr, sb)
			}
			sb.WriteString(": ")
			sb.WriteString(v[0].Error())
		case verrors.Errors:
			sb.WriteString(".")
			buildPathError(v, sb)
		default:
			sb.WriteString(": ")
			sb.WriteString(v.Error())
		}
	}
}
