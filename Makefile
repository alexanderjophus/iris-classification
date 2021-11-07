.PHONY: cmd/predict/models/theta.bin
cmd/predict/models/theta.bin:
	go run cmd/train/main.go
	mv theta.bin cmd/predict/models/theta.bin