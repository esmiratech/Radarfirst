package main

import (
	"context"
	"flag"
	"hellogrpc/proto"
	"log/slog"
	"math"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.String("port", "", "port of grpc server e.g., :5050")

	flag.Parse()

	ctx := context.Background()

	if strings.TrimSpace(*port) == "" {
		slog.Error("port is required")
		os.Exit(1)
	}

	cfg, err := pgxpool.ParseConfig("")
	if err != nil {
		slog.Error("unable to parse config", slog.Any("err", err))
		os.Exit(1)
	}

	conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		slog.Error("unable to connect to database", slog.Any("err", err))
		os.Exit(1)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(ChainStreamInterceptors(
			StreamLogInterceptor,
		)),
		grpc.UnaryInterceptor(ChainUnaryInterceptors(
			UnaryLogInterceptor,
		)),

		grpc.MaxRecvMsgSize(math.MaxInt32),
		grpc.MaxSendMsgSize(math.MaxInt32),
	)

	proto.RegisterGreeterServer(s, &helloService{db: conn})
	grpc_health_v1.RegisterHealthServer(s, &healthServer{db: conn})

	reflection.Register(s)

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		slog.Error("tcp listen", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Info("grpc server started", slog.Any("port", *port))

	if err := s.Serve(lis); err != nil {
		slog.Error("grpc serve", slog.Any("err", err))
		os.Exit(1)
	}
}
