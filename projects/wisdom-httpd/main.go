package main

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"x-learn/projects/wisdom-httpd/handler"
)

func (t *EchoTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// flag
var (
	configFile = pflag.StringP("conf", "c", "./config.yaml", "Application configuration YAML file name")
	debug      = pflag.BoolP("debug", "d", false, "Enable debug mode")
)

func main() {
	// flag
	if err := parseInputFlagConfig(); err != nil {
		log.Fatalf("init config got err: %s", err)
	}

	// service
	e := echo.New()

	// middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           GetLogFormat(),
		CustomTimeFormat: GetLogTimeFormat(),
	}))

	// 模板渲染
	e.Renderer = GetRenderTemplate()

	// static
	e.Static("/", GetAssetPath())

	// handler
	e.GET("/", handler.IndexHandler)
	e.GET("/error", handler.ErrorHandler)

	// handler - wisdom
	e.GET("/wisdom", handler.WisdomHandler)

	// handler - code
	e.GET("/code", handler.CodeHandler)

	// handler - 文件下载
	e.GET("/files/upload", handler.UploadHandler)
	e.GET("/files/download", handler.DownloadHandler)

	// http server
	addr := GetListenAddr()
	log.Debugf("listen: %v", addr)
	e.Logger.Fatal(e.Start(addr))
}

// 解析os参数配置
func parseInputFlagConfig() error {
	pflag.Parse()

	// 应用配置
	if err := ParseConfig(*configFile); err != nil {
		return errors.Wrapf(err, "parse config file %s got err", *configFile)
	}

	// 调试登记
	if *debug {
		log.SetLevel(log.DEBUG)
	}

	return nil
}
