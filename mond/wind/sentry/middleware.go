package sentry

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-errors/errors"
	"github.com/tangbo/twatt/mond/wind/config"
	"github.com/tangbo/twatt/mond/wind/logger"
	"github.com/tangbo/twatt/mond/wind/mq/rabbit/function"
	"github.com/tangbo/twatt/mond/wind/utils/constant"
	mctx "github.com/tangbo/twatt/mond/wind/utils/ctx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func GrpcServerMiddleware() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	timeoutConf := config.GetMethodTimeoutConfig()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		timeout := timeoutConf.GetTimeout(fmt.Sprintf("server.%s", info.FullMethod))
		//fmt.Println(fmt.Sprintf("server.%s", info.FullMethod))
		//fmt.Println(timeout)
		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*timeout)
		defer cancel()
		defer func() {
			if e := recover(); e != nil {
				stackInfo := string(debug.Stack())
				event := sentry.NewEvent()
				err = errors.New(e)
				//event.Message = err.Error() + "\n" + stackInfo
				event.Exception = []sentry.Exception{
					sentry.Exception{
						Type:       err.Error(),
						Stacktrace: sentry.ExtractStacktrace(err),
					},
				}
				event.ServerName = config.GetAppid()
				event.Environment = config.GetEnv()
				event.Extra["env"] = config.GetEnv()
				event.Extra["appId"] = config.GetAppid()
				event.Extra["method"] = info.FullMethod
				event.Extra["traceId"] = mctx.GetTraceId(ctx)
				sentry.CaptureEvent(event)
				logger.GetLogger().Error(ctx, "panic", zap.Any("eventId", event.EventID), zap.Any("stack", stackInfo))
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}

func Consumer(f function.ConsumerFunc) function.ConsumerFunc {
	timeoutConf := config.GetMethodTimeoutConfig()
	return func(ctx context.Context, queue string) (err error) {
		name := fmt.Sprintf("consumer.queue.%s", queue)
		if ctx.Value(constant.AsyncMethodCtxKey) != nil {
			name = fmt.Sprintf("consumer.async.%s", ctx.Value(constant.AsyncMethodCtxKey))
		}
		timeout := timeoutConf.GetTimeout(fmt.Sprintf("server.%s", name))
		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*timeout)
		defer cancel()
		defer func() {
			if e := recover(); e != nil {
				stackInfo := string(debug.Stack())
				event := sentry.NewEvent()
				err = errors.New(e)
				//event.Message = err.Error() + "\n" + stackInfo
				event.Exception = []sentry.Exception{
					sentry.Exception{
						Type:       err.Error(),
						Stacktrace: sentry.ExtractStacktrace(err),
					},
				}
				event.ServerName = config.GetAppid()
				event.Environment = config.GetEnv()
				event.Extra["env"] = config.GetEnv()
				event.Extra["appId"] = config.GetAppid()
				event.Extra["queue"] = queue
				event.Extra["traceId"] = mctx.GetTraceId(ctx)
				sentry.CaptureEvent(event)
				logger.GetLogger().Error(ctx, "panic", zap.Any("eventId", event.EventID), zap.Any("stack", stackInfo))
			}
		}()
		err = f(ctx, queue)
		return err
	}
}

//func GatewaySentryMiddleware(f function2.DispatchFunc, method string) function2.DispatchFunc {
//	return func(ctx context.Context, conn connection.Connection, req proto.Message) (resp proto.Message, err error) {
//		defer func() {
//			if e := recover(); e != nil {
//				stackInfo := string(debug.Stack())
//				event := sentry.NewEvent()
//				err = errors.New(e)
//				//event.Message = err.Error() + "\n" + stackInfo
//				event.Exception = []sentry.Exception{
//					sentry.Exception{
//						Type:       err.Error(),
//						Stacktrace: sentry.ExtractStacktrace(err),
//					},
//				}
//				event.ServerName = config.GetAppid()
//				event.Environment = config.GetEnv()
//				event.Extra["env"] = config.GetEnv()
//				event.Extra["appId"] = config.GetAppid()
//				event.Extra["method"] = method
//				event.Extra["traceId"] = mctx.GetTraceId(ctx)
//				sentry.CaptureEvent(event)
//				logger.GetLogger().Error(ctx, "panic", zap.Any("eventId", event.EventID), zap.Any("stack", stackInfo))
//			}
//		}()
//		resp, err = f(ctx, conn, req)
//		return resp, err
//	}
//}
