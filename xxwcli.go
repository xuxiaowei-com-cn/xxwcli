package main

import (
	"fmt"
	"github.com/xuxiaowei-com-cn/xxwcli/ip"
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
		ip.GetIPCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
