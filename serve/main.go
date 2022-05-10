package main

import (
	"esp_webrtc/commands/daemon"
	"fmt"
	"os"

	"github.com/kardianos/service"
)

func main() {
	d := daemon.NewDaemon()

	s, err := service.New(d, d.Config())
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			x := s.Install()
			if x != nil {
				fmt.Println("error:", x.Error())
				return
			}
			fmt.Println("服务安装成功")
			return
		} else if os.Args[1] == "uninstall" {
			x := s.Uninstall()
			if x != nil {
				fmt.Println("error:", x.Error())
				return
			}
			fmt.Println("服务卸载成功")
			return
		}
	}
	err = s.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
