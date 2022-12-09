package main

import (
	"github.com/urfave/cli"
	"os"
)

var (
	Name        = "xxwcli"
	Version     = "v0.0.1"
	Author      = "xuxiaowei-com-cn/xxwcli: https://github.com/xuxiaowei-com-cn/xxwcli"
	Email       = "徐晓伟 <xuxiaowei@xuxiaowei.com.cn>"
	Copyright   = "徐晓伟"
	Description = "命令行工具"
)

func main() {
	app := &cli.App{}

	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Copyright = Copyright
	app.Description = Description

	err := (app).Run(os.Args)
	if err != nil {
		return
	}

}
