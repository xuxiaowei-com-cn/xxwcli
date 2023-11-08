package port

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
	"sync"
	"time"
)

// GetPortCommand 返回扫描开放端口的 CLI 命令
func GetPortCommand() *cli.Command {
	return &cli.Command{
		Name:  "port",
		Usage: "扫描开放端口",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "127.0.0.1",
				Usage: "扫描主机地址",
			},
			&cli.IntFlag{
				Name:  "startPort",
				Value: 1,
				Usage: "开始端口",
			},
			&cli.IntFlag{
				Name:  "endPort",
				Value: 65535,
				Usage: "结束端口",
			},
		},
		Action: func(c *cli.Context) error {

			host := c.String("host")
			startPort := c.Int("startPort")
			endPort := c.Int("endPort")

			fmt.Printf("主机地址 %s\n", host)
			fmt.Printf("开始端口 %d\n", startPort)
			fmt.Printf("结束端口 %d\n", endPort)
			fmt.Println("开放端口:")

			// 定义 WaitGroup 对象
			var wg sync.WaitGroup

			// 遍历端口范围，并依次创建 goroutine 进行连接尝试
			for port := startPort; port < endPort; port++ {
				wg.Add(1)
				go task(host, port, &wg)

				// 判断是否达到最大并发数量，如果达到则等待一段时间
				if (port+1)%endPort == 0 {
					fmt.Println("Waiting for Goroutines to finish...")
					time.Sleep(time.Second)
				}
			}

			// 等待所有 goroutine 执行完毕
			wg.Wait()

			fmt.Println("All Goroutines Done!")

			return nil
		},
	}
}

func task(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err == nil {
		fmt.Printf("%d\n", port)
		err := conn.Close()
		if err != nil {
			return
		}
	}
}
