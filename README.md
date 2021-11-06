# Iris Classification

This is an example service using gRPC, machine learning, and a bunch of other cool tech.

To find more details, please look at [my blog](https://trelore.github.io/)!

## cmd utilities

The utilities in `cmd` are used to play and experiment with the data outside of the gRPC service.

To run the training utility;

`go run cmd/train/main.go`

This spits out a classifier `model.cls`, which we can then use in the prediction utility;

`go run cmd/predict/main.go`