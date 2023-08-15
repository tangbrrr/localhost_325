package mgrpc

import (
	"github.com/tangbo/twatt/mond/wind/config"
	merr "github.com/tangbo/twatt/mond/wind/err"
	"github.com/tangbo/twatt/mond/wind/logger"
	"github.com/tangbo/twatt/mond/wind/sentinel/breaker"
	"github.com/tangbo/twatt/mond/wind/trace"
	"google.golang.org/grpc"
)

func ClientMiddleware(opt config.GrpcClientOption) []grpc.UnaryClientInterceptor {
	item := []grpc.UnaryClientInterceptor{
		trace.GrpcClientMiddleware(),
	}
	if opt.OpenLog {
		item = append(item, logger.GrpcClientMiddleware)
	}
	item = append(item, merr.GrpcClientMiddleware)
	if config.GetSentinelOption().Breaker.GrpcClientOpen {
		item = append(item, breaker.MakeGRpcClientMiddleware())
	}

	return item
}
