package handler

import (
	"context"
	"github.com/tangbo/twatt/mond/service/biz.demo/app"
	"github.com/tangbo/twatt/mond/service/biz.demo/proto"
	"github.com/tangbo/twatt/mond/wind/logger"
)

type BizdemoService struct {
	_log logger.Logger
	app  *app.App
}

func (m *BizdemoService) Ping(ctx context.Context, n *Bizdemo.PingReq) (*Bizdemo.PingResp, error) {
	panic("implement me")
}

func NewBizdemoService() *BizdemoService {
	return &BizdemoService{_log: logger.GetLogger()}
}
