package main

import (
	"google.golang.org/grpc/codes"
	"namespacelabs.dev/foundation/framework/rpcerrors"
)

func WrapError(errMsg error) error {
	return rpcerrors.Errorf(codes.Internal, "%w", errMsg)
}
