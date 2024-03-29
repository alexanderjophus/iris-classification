// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: iris_classification/v1/service.proto

package irisclassificationpbconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/trelore/iris-classification/proto/gen/go/iris_classification/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// IrisClassificationServiceName is the fully-qualified name of the IrisClassificationService
	// service.
	IrisClassificationServiceName = "iris_classification.v1.IrisClassificationService"
)

// IrisClassificationServiceClient is a client for the
// iris_classification.v1.IrisClassificationService service.
type IrisClassificationServiceClient interface {
	// Predict the Iris Classification
	Predict(context.Context, *connect_go.Request[v1.PredictRequest]) (*connect_go.Response[v1.PredictResponse], error)
}

// NewIrisClassificationServiceClient constructs a client for the
// iris_classification.v1.IrisClassificationService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIrisClassificationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IrisClassificationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &irisClassificationServiceClient{
		predict: connect_go.NewClient[v1.PredictRequest, v1.PredictResponse](
			httpClient,
			baseURL+"/iris_classification.v1.IrisClassificationService/Predict",
			opts...,
		),
	}
}

// irisClassificationServiceClient implements IrisClassificationServiceClient.
type irisClassificationServiceClient struct {
	predict *connect_go.Client[v1.PredictRequest, v1.PredictResponse]
}

// Predict calls iris_classification.v1.IrisClassificationService.Predict.
func (c *irisClassificationServiceClient) Predict(ctx context.Context, req *connect_go.Request[v1.PredictRequest]) (*connect_go.Response[v1.PredictResponse], error) {
	return c.predict.CallUnary(ctx, req)
}

// IrisClassificationServiceHandler is an implementation of the
// iris_classification.v1.IrisClassificationService service.
type IrisClassificationServiceHandler interface {
	// Predict the Iris Classification
	Predict(context.Context, *connect_go.Request[v1.PredictRequest]) (*connect_go.Response[v1.PredictResponse], error)
}

// NewIrisClassificationServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIrisClassificationServiceHandler(svc IrisClassificationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/iris_classification.v1.IrisClassificationService/Predict", connect_go.NewUnaryHandler(
		"/iris_classification.v1.IrisClassificationService/Predict",
		svc.Predict,
		opts...,
	))
	return "/iris_classification.v1.IrisClassificationService/", mux
}

// UnimplementedIrisClassificationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIrisClassificationServiceHandler struct{}

func (UnimplementedIrisClassificationServiceHandler) Predict(context.Context, *connect_go.Request[v1.PredictRequest]) (*connect_go.Response[v1.PredictResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("iris_classification.v1.IrisClassificationService.Predict is not implemented"))
}
