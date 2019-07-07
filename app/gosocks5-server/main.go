package main

// Socks5其实很简单, 任何人都可以参考规范很简单的写出一个sock5代理

import (
	"fmt"
	"log"
	"net"

	"github.com/xiangrui2019/go_socks5"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// 默认配置
	config := &go_socks5.Config{
		ListenAddr: fmt.Sprintf(":%d", 7448),
		// 密码随机生成
		Password: go_socks5.RandPassword(),
	}
	config.ReadConfig()
	config.SaveConfig()

	// 启动 server 端并监听
	lsServer, err := go_socks5.NewRemoteServer(config.Password, config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsServer.Listen(func(listenAddr net.Addr) {
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
密码 password：
%s
	`, listenAddr, config.Password))
		log.Printf("gosocks5-server 启动成功 监听在 %s\n", listenAddr.String())
	}))
}
