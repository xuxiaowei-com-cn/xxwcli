package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/xuxiaowei-com-cn/xxwcli/ip"
	"github.com/xuxiaowei-com-cn/xxwcli/web"
	"os"

	"github.com/urfave/cli"
)

var (
	Name        = "xxwcli"
	Version     = "v0.0.1"
	Author      = "xuxiaowei-com-cn/xxwcli: https://github.com/xuxiaowei-com-cn/xxwcli"
	Email       = "徐晓伟 <xuxiaowei@xuxiaowei.com.cn>"
	Copyright   = "徐晓伟工作室 <xuxiaowei@xuxiaowei.com.cn>"
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
			Name:      "version",
			ShortName: "v",
			Usage:     "获取版本信息",
			Action: func(c *cli.Context) error {
				fmt.Printf("name %s\n", Name)
				fmt.Printf("version %s\n", Version)
				fmt.Printf("author %s\n", Author)
				fmt.Printf("email %s\n", Email)
				fmt.Printf("copyright %s\n", Copyright)
				fmt.Printf("description %s\n", Description)

				fmt.Printf("buildTime %s\n", buildTime)
				fmt.Printf("commitSha %s\n", commitSha)
				fmt.Printf("commitShortSha %s\n", commitShortSha)
				return nil
			},
		},
		ip.GetIPCommand(),
		web.GetHttpdCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(red(err))
		os.Exit(1)
	}
}

var (
	// buildTime 表示构建时间。
	buildTime string
	// 项目为其构建的提交修订
	commitSha string
	// CI_COMMIT_SHA 的前八个字符
	commitShortSha string
)

func init() {
	if buildTime == "" {
		buildTime = "unknown"
	}
	if commitSha == "" {
		commitSha = "unknown"
	}
	if commitShortSha == "" {
		commitShortSha = "unknown"
	}
}
