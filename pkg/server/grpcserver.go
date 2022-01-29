package server

import (
	"context"
	"getputer/internal/cache"
	api "getputer/pkg/api/versions/v1"
	v1 "getputer/pkg/api/versions/v1"
)

type GRPCServer struct {
	v1.UnimplementedGet_PutServer
}

var c = cache.New()

func (s *GRPCServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	return &api.GetResponse{Value: c.Get(req.Name)}, nil
}

func (s *GRPCServer) Put(ctx context.Context, req *api.PutRequest) (*api.PutResponse, error) {
	c.Put(req.Name, req.Value)
	return &api.PutResponse{Name: req.Name, Value: req.Value}, nil
}
