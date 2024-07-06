package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App    AppConfig `yaml:"app"`
	Listen string    `yaml:"listen"` // 监听
}

// AppConfig 应用配置
type AppConfig struct {
	Root          string `yaml:"root"`       // 配置文件
	LogFormat     string `yaml:"log_format"` // 日志格式
	LogTimeFormat string `yaml:"log_time_format"`
	Assets        struct {
		AssetPath string `yaml:"asset_path"` // 静态资源path
		ViewPath  string `yaml:"view_path"`  // 视图资源path
	} `yaml:"assets"`
}

// 默认配置
var defaultConfig *Config

// ParseConfig 解析系统配置
func ParseConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return errors.Wrapf(err, "read filename fail: %v", filename)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return errors.Wrapf(err, "yaml unmarshal config fail")
	}

	defaultConfig = &cfg

	return nil
}

// GetAppConfig 返回系统配置
func GetAppConfig() *AppConfig {
	if defaultConfig != nil {
		return &defaultConfig.App
	}
	return &AppConfig{}
}

// GetRoot 返回项目根目录
func GetRoot() string {
	return GetAppConfig().Root
}

// GetListenAddr 监听地址
func GetListenAddr() string {
	return defaultConfig.Listen
}

// GetLogFormat 获取应用的日志格式
func GetLogFormat() string {
	return GetAppConfig().LogFormat + "\n"
}

// GetLogTimeFormat 日志时间格式
func GetLogTimeFormat() string {
	return GetAppConfig().LogTimeFormat
}

// GetRootPath 获取单个path地址
func GetRootPath(path string) string {
	return fmt.Sprintf("%s/%s", GetRoot(), path)
}

// GetAssetPath 静态资源路径
func GetAssetPath() string {
	return fmt.Sprintf("%s/%s", GetRoot(), GetAppConfig().Assets.AssetPath)
}

// GetSpecialViewPath 视图地址, assets/views/path
func GetSpecialViewPath(path string) string {
	viewPath := GetAppConfig().Assets.ViewPath
	return fmt.Sprintf("%s/%s/%s", GetRoot(), viewPath, path)
}

// GetSpecialViewPathList 获取一批path地址
func GetSpecialViewPathList(paths ...string) []string {
	var ret []string
	for _, path := range paths {
		ret = append(ret, GetSpecialViewPath(path))
	}

	return ret
}
