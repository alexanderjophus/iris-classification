package server

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"math"

	"github.com/bufbuild/connect-go"
	pb "github.com/trelore/iris-classification/proto/gen/go/iris_classification/v1"
	"github.com/trelore/iris-classification/svc/models"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// New returns a new S
func New() *S {
	var thetaT *tensor.Dense
	err := gob.NewDecoder(bytes.NewReader(models.Data)).Decode(&thetaT)
	if err != nil {
		log.Fatal(err)
	}
	return &S{thetaT: thetaT}
}

// S Implements the IrisClassificationService service
type S struct {
	thetaT *tensor.Dense
}

// Predict implements proto
func (s *S) Predict(ctx context.Context, req *connect.Request[pb.PredictRequest]) (*connect.Response[pb.PredictResponse], error) {
	g := gorgonia.NewGraph()
	theta := gorgonia.NodeFromAny(g, s.thetaT, gorgonia.WithName("theta"))

	values := make([]float64, 5)
	xT := tensor.New(tensor.WithBacking(values))
	x := gorgonia.NodeFromAny(g, xT, gorgonia.WithName("x"))
	y, err := gorgonia.Mul(x, theta)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()

	values[4] = 1.0
	values[0] = req.Msg.GetSepalLength()
	values[1] = req.Msg.GetSepalWidth()
	values[2] = req.Msg.GetPetalLength()
	values[3] = req.Msg.GetPetalWidth()

	if err = machine.RunAll(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
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
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("unknown iris"))
	}
	res := connect.NewResponse(&pb.PredictResponse{
		Predicition: class,
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
