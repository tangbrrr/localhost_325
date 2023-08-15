package pool

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/getsentry/sentry-go"
	"github.com/tangbo/twatt/mond/wind/config"
	"github.com/tangbo/twatt/mond/wind/logger"
	"go.uber.org/zap"
)

func SafeGo(f func()) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				stackInfo := string(debug.Stack())
				event := sentry.NewEvent()
				//event.Message = stackInfo
				err := errors.New(fmt.Sprintf("%v", e))
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
				sentry.CaptureEvent(event)
				logger.GetLogger().Error(context.TODO(), "panic", zap.Any("eventId", event.EventID), zap.Any("stack", stackInfo))
			}
		}()
		f()
	}()
}
