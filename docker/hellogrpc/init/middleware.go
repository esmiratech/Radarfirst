package main

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

// StreamLogInterceptor writes to the log before and after calling handler and logs any errors returned from handler.
func StreamLogInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()

	slog.Info("started", slog.Any("source", "stream"), slog.Any("full_method", info.FullMethod))

	defer slog.Info("completed", slog.Any("source", "stream"), slog.Any("full_method", info.FullMethod), slog.Any("duration", time.Since(start)))

	err := handler(srv, stream)
	if err != nil {
		slog.Error("error", slog.Any("source", "stream"), slog.Any("full_method", info.FullMethod), slog.Any("err", err))
	}

	return err
}

// UnaryLogInterceptor writes to the log before and after calling handler and
// logs any errors returned from handler.
func UnaryLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	slog.Info("started", slog.Any("source", "unary"), slog.Any("full_method", info.FullMethod))

	defer slog.Info("completed", slog.Any("source", "unary"), slog.Any("full_method", info.FullMethod), slog.Any("duration", time.Since(start)))

	v, err := handler(ctx, req)
	if err != nil {
		slog.Error("error", slog.Any("source", "unary"), slog.Any("full_method", info.FullMethod), slog.Any("err", err))
	}

	return v, err
}

// ChainUnaryInterceptors returns a UnaryServerInterceptor that calls each of
// the provided interceptors in order, starting with the leftmost.
func ChainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		buildChain := func(h grpc.UnaryServerInterceptor, next grpc.UnaryHandler) grpc.UnaryHandler {
			return func(ctx context.Context, req interface{}) (interface{}, error) {
				return h(ctx, req, info, next)
			}
		}
		next := handler

		for i := len(interceptors) - 1; i >= 0; i-- {
			next = buildChain(interceptors[i], next)
		}

		return next(ctx, req)
	}
}

// ChainStreamInterceptors returns a StreamServerInterceptor that calls each of
// the provided interceptors in order, starting with the leftmost.
func ChainStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		buildChain := func(h grpc.StreamServerInterceptor, next grpc.StreamHandler) grpc.StreamHandler {
			return func(srv interface{}, stream grpc.ServerStream) error {
				return h(srv, stream, info, next)
			}
		}
		next := handler

		for i := len(interceptors) - 1; i >= 0; i-- {
			next = buildChain(interceptors[i], next)
		}

		return next(srv, stream)
	}
}
