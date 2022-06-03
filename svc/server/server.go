package server

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"math"

	pb "github.com/trelore/iris-classification/proto/gen/go/iris_classification/v1"
	"github.com/trelore/iris-classification/svc/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// New returns a new S
func New() S {
	var thetaT *tensor.Dense
	err := gob.NewDecoder(bytes.NewReader(models.Data)).Decode(&thetaT)
	if err != nil {
		log.Fatal(err)
	}
	return S{thetaT: thetaT}
}

// S Implements the IrisClassificationService service
type S struct {
	thetaT *tensor.Dense
	pb.UnimplementedIrisClassificationServiceServer
}

// Predict implements proto
func (s *S) Predict(ctx context.Context, req *pb.PredictRequest) (*pb.PredictResponse, error) {
	g := gorgonia.NewGraph()
	theta := gorgonia.NodeFromAny(g, s.thetaT, gorgonia.WithName("theta"))

	values := make([]float64, 5)
	xT := tensor.New(tensor.WithBacking(values))
	x := gorgonia.NodeFromAny(g, xT, gorgonia.WithName("x"))
	y, err := gorgonia.Mul(x, theta)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()

	values[4] = 1.0
	values[0] = req.GetSepalLength()
	values[1] = req.GetSepalWidth()
	values[2] = req.GetPetalLength()
	values[3] = req.GetPetalWidth()

	if err = machine.RunAll(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer machine.Reset()

	var class string
	switch math.Round(y.Value().Data().(float64)) {
	case 1:
		class = "setosa"
	case 2:
		class = "versicolor"
	case 3:
		class = "virginica"
	default:
		return nil, status.Error(codes.Internal, "unknown iris")
	}
	return &pb.PredictResponse{
		Predicition: class,
	}, nil
}
