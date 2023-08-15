package sentry

import (
	"sync"

	"github.com/getsentry/sentry-go"
	"github.com/tangbo/twatt/mond/wind/config"
)

var (
	sentryOnce sync.Once
)

func Init() {
	sentryOnce.Do(func() {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:         config.GetSentry(),
			ServerName:  config.GetAppid(),
			Environment: config.GetEnv(),
		})
		if err != nil {
			panic(err)
		}
	})
}
