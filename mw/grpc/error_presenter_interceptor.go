package zerovalidationgrpcmw

import (
	"context"
	"errors"
	verrors "github.com/0B1t322/zero-validation/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorPresenter ...
type ErrorPresenter interface {
	Present(errors verrors.Errors) error
}

// WithErrorPresenter ...
func WithErrorPresenter(presenter ErrorPresenter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)

		return resp, handlePresenterError(presenter, err)
	}
}

func handlePresenterError(presenter ErrorPresenter, err error) error {
	if err == nil {
		return nil
	}

	var verr verrors.Errors
	if !errors.Is(err, verr) {
		return err
	}

	if !errors.As(err, &verr) {
		return err
	}

	presentedErr := presenter.Present(verr)
	return status.Error(codes.InvalidArgument, presentedErr.Error())
}
