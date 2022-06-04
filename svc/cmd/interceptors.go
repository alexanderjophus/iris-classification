package cmd

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func NewLoggingInterceptor(logger *zap.Logger) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			name := req.Spec().Procedure
			start := time.Now()
			res, err := next(ctx, req)
			if err != nil {
				return nil, err
			}
			logger.Sugar().Info(
				zap.String("name", name),
				zap.Duration("duration", time.Since(start)),
				zap.Any("info", req.Any()),
			)
			return res, nil
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func NewPromInterceptor(counter *prometheus.CounterVec) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			if err != nil {
				return nil, err
			}
			counter.With(prometheus.Labels{
				"rpc": req.Spec().Procedure,
			})
			return res, nil
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
