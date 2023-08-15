package main

import (
	"github.com/tangbo/twatt/mond/service/biz.demo/handler"
	"github.com/tangbo/twatt/mond/wind"
)

func main() {
	wind.InitFrame(handler.NewHook())
}
