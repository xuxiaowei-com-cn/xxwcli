package web

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"net/http"
)

// GetHttpdCommand 返回 httpd 代理的 CLI 命令
func GetHttpdCommand() cli.Command {
	return cli.Command{
		Name:  "httpd",
		Usage: "httpd 代理",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dir",
				Value: ".",
				Usage: "指定要提供代理的目录",
			},
			cli.StringFlag{
				Name:  "port",
				Value: "8080",
				Usage: "指定代理服务的端口",
			},
		},
		Action: func(c *cli.Context) error {
			dir := c.String("dir")
			port := c.String("port")

			msg := fmt.Sprintf("代理服务端口：%s\n", port)
			msg += fmt.Sprintf("代理服务目录：%s\n", dir)

			green := color.New(color.FgGreen).SprintFunc()

			fmt.Print(green(msg))

			http.Handle("/", http.FileServer(http.Dir(dir)))
			err := http.ListenAndServe(":"+port, nil)

			if err != nil {
				return err
			}
			return nil
		},
	}
}
