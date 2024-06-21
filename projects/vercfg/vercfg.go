package vercfg

import (
	"log"

	"github.com/pkg/errors"
)

type VerCfg interface {
	// GetVerConfig 获取配置
	GetVerConfig() (interface{}, error)

	// SetConfig 设置配置
	SetConfig(ver string, cfg interface{}) error
}

// 获取指定版本配置
func getConfig(ver string) interface{} {
	return nil
}

// 每个版本配置控制
type vc map[string]interface{}

func init() {
	// 新增一些配置
	appcfgs := map[string]vc{
		"ios": {
			"5.4":    "cfg_5.4",
			"6.4.24": "cfg_6.4.24",
			"6.4.3":  "cfg_6.4.3",
			"6.5.24": "cfg_6.5.24",
			"7.4.24": "cfg_7.4.24",
		},
	}

	// 初始化ios配置
	appInfo := &App{
		Platform:   "ios",
		NowVersion: "7.4.24",
	}

	for _, vc := range appcfgs {
		for ver, cfg := range vc {
			if err := appInfo.SetConfig(ver, cfg); err != nil {
				log.Fatalf("app set fail: %s", err)
			}
		}
	}
}

func GetNearVersionCfg(ver string) interface{} {
	appInfo := &App{
		Platform:   "ios",
		NowVersion: "7.4.24",
	}

	// 新增以
	cfg, err := appInfo.GetVerConfig()
	if err != nil {
		return errors.Wrap(err, "get version cfg got err")
	}
	return cfg
}

type storage struct {
}

// App App版本控制
type App struct {
	Platform   string
	NowVersion string
}

func (app *App) GetVerConfig() (interface{}, error) {
	// TODO implement me
	panic("implement me")
}

func (app *App) SetConfig(ver string, cfg interface{}) error {
	// TODO implement me
	panic("implement me")
}
