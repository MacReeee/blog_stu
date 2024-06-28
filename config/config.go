package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName    string
	Version    float32
	CurrentDir string

	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	Cfg = &tomlConfig{}
	Cfg.System.AppName = "go-blog-practice"
	Cfg.System.Version = 1.0
	curDir, _ := os.Getwd()
	Cfg.System.CurrentDir = curDir
	if _, err := toml.DecodeFile("config/config.toml", &Cfg); err != nil {
		panic(err)
	}
}
