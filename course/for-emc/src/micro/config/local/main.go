package main

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 加载配置文件
	if err := config.Load(file.NewSource(
		file.WithPath("./config/config.yml"),
		file.WithPath("./config/config.json"),
	)); err != nil {
		log.Error(err)
		return
	}

	// 根据实际情况，定义合适的结构
	// go-config通过scan方法将配置转成JSON，再传入指定类型的field中
	type Host struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	var host Host
	if err := config.Get("hosts", "database").Scan(&host); err != nil {
		log.Error(err)
		return
	}

	log.Info(host.Name, host.Address, host.Port)
}
