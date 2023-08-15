package svc

import (
	"github.com/tangbo/twatt/liyue/im/gateway/internal/config"
	"github.com/tangbo/twatt/liyue/im/rpc/im"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	ImRpc  im.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ImRpc:  im.NewIm(zrpc.MustNewClient(c.ImRpc)),
	}
}
