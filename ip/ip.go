package ip

import (
	"fmt"
	"github.com/urfave/cli"
	"net"
	"os"
)

// GetIPCommand 返回获取 IP 地址的 CLI 命令
func GetIPCommand() cli.Command {
	return cli.Command{
		Name:  "ip",
		Usage: "获取当前计算机的 IP 地址",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "v4",
				Usage: "仅显示 IPv4 地址",
			},
			&cli.BoolFlag{
				Name:  "v6",
				Usage: "仅显示 IPv6 地址",
			},
		},
		Action: func(c *cli.Context) error {
			addrs, err := net.InterfaceAddrs()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			ips := make(map[string]bool)
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					var isIPv4 = ipnet.IP.To4() != nil
					var isIPv6 = ipnet.IP.To16() != nil && ipnet.IP.To4() == nil

					if (c.Bool("v4") && isIPv4) || (c.Bool("v6") && isIPv6) || (!c.Bool("v4") && !c.Bool("v6")) {
						if _, ok := ips[ipnet.IP.String()]; !ok {
							ips[ipnet.IP.String()] = true
							fmt.Println(ipnet.IP.String())
						}
					}
				}
			}
			return nil
		},
	}
}
