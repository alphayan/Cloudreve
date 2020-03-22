package main

import (
	"flag"

	"github.com/HFO4/cloudreve/bootstrap"
	"github.com/HFO4/cloudreve/pkg/conf"
	"github.com/HFO4/cloudreve/pkg/util"
	"github.com/HFO4/cloudreve/routers"
)

var confPath string

func init() {
	flag.StringVar(&confPath, "c", util.RelativePath("conf.ini"), "配置文件路径")
	flag.Parse()
	bootstrap.Init(confPath)
}

func main() {
	api := routers.InitRouter()
	util.Log().Info("开始监听 %s", conf.SystemConfig.Listen)
	if err := api.Run(conf.SystemConfig.Listen); err != nil {
		util.Log().Error("无法监听[%s]，%s", conf.SystemConfig.Listen, err)
	}
}
