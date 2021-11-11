FROM golang:1.17 AS builder
WORKDIR /go/src/github.com/trelore/iris-classification/
COPY . /go/src/github.com/trelore/iris-classification
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./svc

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/trelore/iris-classification/app ./
CMD ["./app"]