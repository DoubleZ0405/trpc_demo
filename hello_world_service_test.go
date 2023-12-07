package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/DoubleZ0405/trpc_demo/helloworld"
	"github.com/golang/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/github.com/DoubleZ0405/trpc_demo/helloworld/helloword_mock.go -package=helloworld -self_package=github.com/DoubleZ0405/trpc_demo/helloworld --source=stub/github.com/DoubleZ0405/trpc_demo/helloworld/helloword.trpc.go

func Test_helloWorldServiceImpl_Hello(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helloWorldServiceService := pb.NewMockHelloWorldServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := helloWorldServiceService.EXPECT().Hello(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
		s := &helloWorldServiceImpl{}
		return s.Hello(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.HelloRequest
		rsp *pb.HelloResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.HelloResponse
			var err error
			if rsp, err = helloWorldServiceService.Hello(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("helloWorldServiceImpl.Hello() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("helloWorldServiceImpl.Hello() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
