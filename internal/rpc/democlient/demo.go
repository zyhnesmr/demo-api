// Code generated by goctl. DO NOT EDIT.
// Source: mydemo.proto

package democlient

import (
	"context"

	"demo-api/internal/rpc/demo"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InMsg  = demo.InMsg
	OutMsg = demo.OutMsg

	Demo interface {
		SayHello(ctx context.Context, in *InMsg, opts ...grpc.CallOption) (*OutMsg, error)
	}

	defaultDemo struct {
		cli zrpc.Client
	}
)

func NewDemo(cli zrpc.Client) Demo {
	return &defaultDemo{
		cli: cli,
	}
}

func (m *defaultDemo) SayHello(ctx context.Context, in *InMsg, opts ...grpc.CallOption) (*OutMsg, error) {
	client := demo.NewDemoClient(m.cli.Conn())
	return client.SayHello(ctx, in, opts...)
}
