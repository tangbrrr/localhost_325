package handler

import (
	"context"
	"fmt"
	"github.com/tangbo/twatt/mond/service/biz.demo/proto"
	"github.com/tangbo/twatt/mond/wind"
	"os"
	"testing"
)

var (
	ctx context.Context
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	h := NewHook()
	wind.InitFrame(h, wind.WithTestMode())
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestBizdemoService_Ping(t *testing.T) {
	req := &Bizdemo.PingReq{}
	resp, err := _handler.Ping(ctx, req)
	t.Error(err)
	fmt.Println(resp)
}
