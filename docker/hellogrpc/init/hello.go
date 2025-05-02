package main

import (
	"context"
	"fmt"
	"hellogrpc/proto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type helloService struct {
	proto.UnimplementedGreeterServer
	db *pgxpool.Pool
}

func (s *helloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	_, err := s.db.Exec(ctx, `INSERT INTO hello.world (name) VALUES ($1)`, req.Name)
	if err != nil {
		fmt.Errorf("insert hello world: %w", err)
	}
	return &proto.HelloReply{Message: fmt.Sprintf("G'day %s", req.Name)}, nil
}
