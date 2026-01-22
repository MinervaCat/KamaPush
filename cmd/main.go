package main

import (
	"KamaPush/internal/config"
	"KamaPush/internal/https_server"
	"KamaPush/internal/service/kafka"
	"KamaPush/internal/service/push"
	"KamaPush/pkg/zlog"
	"fmt"
)

func main() {
	zlog.Info("push服务开始")
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port
	go func() {
		// Ubuntu22.04云服务器部署
		if err := https_server.GE.RunTLS(fmt.Sprintf("%s:%d", host, port), "/etc/ssl/certs/server.crt", "/etc/ssl/private/server.key"); err != nil {
			zlog.Fatal("server running fault")
			return
		}
	}()
	kafka.KafkaService.KafkaInit2()
	push.Pusher.Start()
}
