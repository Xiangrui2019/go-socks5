package main

import (
	"fmt"
	"log"
	"net"

	"github.com/xiangrui2019/go_socks5"
)

func main() {
	log.SetFlags(log.Lshortfile)

	config := &go_socks5.Config{}
	config.ReadConfig()
	config.SaveConfig()

	localserver, err := go_socks5.NewLocalServer(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(localserver.Listen(func(listenAddr net.Addr) {
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
远程服务地址 remote：
%s
密码 password：
%s
	`, listenAddr, config.RemoteAddr, config.Password))
		log.Printf("gosocks5 启动成功 监听在 %s\n", listenAddr.String())
	}))
}
