package cmd

import "text/template"

var cmdTemplate, _ = template.New("").Parse(`
package main

import (
	"github.com/tangbo/twatt/mond/wind"
	"github.com/tangbo/twatt/mond/service/{{.FolderPath}}/handler"
)

func main() {
	wind.InitFrame(handler.NewHook())
}

`)
