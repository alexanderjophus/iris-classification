package server

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/trelore/iris-classification/proto/gen/go/iris_classification/v1"
)

// New returns a new S
func New() S {
	return S{}
}

// S Implements the Confidence service
type S struct {
	pb.UnimplementedIrisClassificationServiceServer
}

var irisFlowers = []string{"iris setosa", "iris versicolor", "iris virginica"}

// Predict implements proto
func (s *S) Predict(ctx context.Context, req *pb.PredictRequest) (*pb.PredictResponse, error) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	iris := irisFlowers[rand.Intn(len(irisFlowers))]

	return &pb.PredictResponse{
		Predicition: iris,
	}, nil
}
