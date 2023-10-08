package web

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
)

// GetHttpdCommand 返回 httpd 代理的 CLI 命令
func GetHttpdCommand() *cli.Command {
	return &cli.Command{
		Name:  "httpd",
		Usage: "http 代理",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Value: ".",
				Usage: "指定提供代理的文件夹",
			},
			&cli.BoolFlag{
				Name:  "dir",
				Usage: "代理目录",
			},
			&cli.IntFlag{
				Name:  "port",
				Value: 8080,
				Usage: "指定代理服务的端口",
			},
		},
		Action: func(c *cli.Context) error {
			dir := c.Bool("dir")
			path := c.String("path")
			port := c.String("port")

			msg := fmt.Sprintf("代理服务端口：%s\n", port)
			msg += fmt.Sprintf("代理服务目录：%s\n", path)

			green := color.New(color.FgGreen).SprintFunc()

			fmt.Print(green(msg))

			http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

				//fmt.Printf("路径：%s，URL：%s\n", path+req.URL.Path, req.URL.Path)

				// 判断是否为主页
				if isHomepage(path, w, req, dir) {
					return
				}

				if req.URL.Path[len(req.URL.Path)-1] == '/' {
					// 判断是否存在 index.html 文件
					if isHomepage(path, w, req, dir) {
						return
					}
				}

				// 默认处理逻辑
				http.FileServer(http.Dir(path)).ServeHTTP(w, req)
			})

			err := http.ListenAndServe(":"+port, nil)

			if err != nil {
				return err
			}
			return nil
		},
	}
}

// 判断请求的路径是否为主页，如果是则检查是否存在 index.html 文件
func isHomepage(path string, w http.ResponseWriter, req *http.Request, dir bool) bool {
	if req.URL.Path == "/" {
		_, err := os.Stat(path + req.URL.Path + "index.html")
		if err != nil {
			// 主页不存在
			if os.IsNotExist(err) {
				// 未开启代理目录
				if !dir {
					// 拒绝返回目录信息
					http.Error(w, "目录下不存在 index.html，不允许访问目录", http.StatusForbidden)
					return true
				}
			}
		}

		http.FileServer(http.Dir(path)).ServeHTTP(w, req)
		return true
	}
	return false
}
