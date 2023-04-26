package main

import (
	"fmt"
	"net"
	"os"

	"github.com/urfave/cli"
)

var (
	Name        = "xxwcli"
	Version     = "v0.0.1"
	Author      = "xuxiaowei-com-cn/xxwcli: https://github.com/xuxiaowei-com-cn/xxwcli"
	Email       = "徐晓伟 <xuxiaowei@xuxiaowei.com.cn>"
	Copyright   = "徐晓伟 <xuxiaowei@xuxiaowei.com.cn>"
	Description = "命令行工具"
)

func main() {
	app := &cli.App{}

	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Usage = Description
	app.Copyright = Copyright

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "获取当前计算机的 IP 地址",
			Action: func(c *cli.Context) error {
				addrs, err := net.InterfaceAddrs()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
						if ipnet.IP.To4() != nil {
							fmt.Println(ipnet.IP.String())
						} else if ipnet.IP.To16() != nil {
							fmt.Println(ipnet.IP.String())
						}
					}
				}
				return nil
			},
		},

		{
			Name:  "ipv4",
			Usage: "获取当前计算机的 IPv4 地址",
			Action: func(c *cli.Context) error {
				addrs, err := net.InterfaceAddrs()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())
					}
				}
				return nil
			},
		},

		{
			Name:  "ipv6",
			Usage: "获取当前计算机的 IPv6 地址",
			Action: func(c *cli.Context) error {
				addrs, err := net.InterfaceAddrs()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To16() != nil && ipnet.IP.To4() == nil {
						fmt.Println(ipnet.IP.String())
					}
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
