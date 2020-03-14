package server

import (
	"context"

	pb "github.com/trelore/iris-classification/proto/gen/go/iris-classification"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// New returns a new S
func New() S {
	return S{}
}

// S Implements the Confidence service
type S struct {
}

// Predict implements proto
func (s S) Predict(ctx context.Context, req *pb.PredictRequest) (*pb.PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confidence not implemented")
}
