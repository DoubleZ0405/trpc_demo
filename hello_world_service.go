package main

import (
	"context"

	pb "github.com/DoubleZ0405/trpc_demo/helloworld"
)

type helloWorldServiceImpl struct {
	pb.UnimplementedHelloWorldService
}

// Hello Hello says hello.
func (s *helloWorldServiceImpl) Hello(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloResponse, error) {
	rsp := &pb.HelloResponse{}
	return rsp, nil
}
