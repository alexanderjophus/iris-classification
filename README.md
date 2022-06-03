# Iris Classification

This is an example service using gRPC, machine learning, and a bunch of other cool tech.

To find more details, please look at [my blog](https://trelore.github.io/tags/iris/)!

## Running

All you need is Go (I think)!

To run this repo, run `make`, this will create a model `theta.bin` (naming is hard), and copy that model over to both the predict cmd and the predict svc.

After that, if you run either `go run cmd/predict/main.go` to run the model against some currently hardcoded values

_or_

Better yet, you can run `go run svc/main.go` this will spin up an RPC service, where you can call `Predict` with the 4 values for sepal/petal width/length using a tool like evans or grpcurl.

## cmd utilities

The utilities in `cmd` are used to play and experiment with the data outside of the gRPC service.

To run the training utility;

`go run cmd/train/main.go`

This spits out a classifier `model.cls`, which we can then use in the prediction utility;

`go run cmd/predict/main.go`