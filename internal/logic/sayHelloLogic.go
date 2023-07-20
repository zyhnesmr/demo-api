package logic

import (
	"context"
	"demo-api/internal/rpc/democlient"
	"errors"

	"demo-api/internal/svc"
	"demo-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SayHelloLogic) SayHello() (resp *types.Response, err error) {
	hello, err := l.svcCtx.DemoClient.SayHello(l.ctx, &democlient.InMsg{Msg: "hello"})
	if err != nil {
		return nil, errors.New("bad request from grpc")
	}

	return &types.Response{Message: hello.Msg}, nil
}
