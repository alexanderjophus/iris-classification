.PHONY: all
all: svc/models/theta.bin cmd/predict/models/theta.bin

.PHONY: theta.bin
theta.bin:
	go run cmd/train/main.go

.PHONY: cmd/predict/models/theta.bin
cmd/predict/models/theta.bin: theta.bin
	cp theta.bin cmd/predict/models/theta.bin

.PHONY: svc/models/theta.bin
svc/models/theta.bin: theta.bin
	cp theta.bin svc/models/theta.bin
