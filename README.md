# xxwcli 命令行（脚手架）工具

## 脚手架命令

```shell
D:\GolandProjects\xxwcli>xxwcli.exe -h
NAME:
   xxwcli - 命令行工具

USAGE:
    [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

AUTHOR:
   xuxiaowei-com-cn/xxwcli: https://github.com/xuxiaowei-com-cn/xxwcli <徐晓伟 <xuxiaowei@xuxiaowei.com.cn>>

COMMANDS:
   ip       获取当前计算机的 IP 地址
   httpd    http 代理
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   徐晓伟 <xuxiaowei@xuxiaowei.com.cn>

D:\GolandProjects\xxwcli>
```

```shell
D:\GolandProjects\xxwcli>xxwcli.exe ip -h
NAME:
ip - 获取当前计算机的 IP 地址

USAGE:
ip [command options] [arguments...]

OPTIONS:
--v4  仅显示 IPv4 地址
--v6  仅显示 IPv6 地址


D:\GolandProjects\xxwcli>
```

```shell
D:\GolandProjects\xxwcli>xxwcli.exe httpd -h
NAME:
httpd - http 代理

USAGE:
httpd [command options] [arguments...]

OPTIONS:
--path value  指定提供代理的文件夹 (default: ".")
--dir         代理目录
--port value  指定代理服务的端口 (default: 8080)


D:\GolandProjects\xxwcli>
```

## 开发命令

### get

```shell
go get github.com/urfave/cli
go get github.com/fatih/color
```

### mod

```shell
go mod download
```

```shell
go mod tidy
```

### run

```shell
go run xxwcli.go
```

### build

```shell
go build xxwcli.go
```

```shell
# -s 表示去掉符号表信息
# -w 表示去掉 DWARF 调试信息
go build -ldflags "-w -s" xxwcli.go
```

- Windows PowerShell

```shell
go build -ldflags "-X main.buildTime=$(Get-Date -Format yyyy-MM-dd_HH:mm:ss) -X main.commitSha=$(git rev-parse HEAD) -X main.commitShortSha=$(git rev-parse --short HEAD) " xxwcli.go
```

### require

1. [github.com/urfave/cli](https://github.com/urfave/cli)
    1. [文档](https://cli.urfave.org)
    2. 一个简单、快速、有趣的程序包，用于在Go中构建命令行应用程序

## 批量添加远端仓库地址

<details>
<summary>点击展开</summary>
git remote add gitee https://gitee.com/xuxiaowei-com-cn/xxwcli.git

git remote add gitlab https://gitlab.com/xuxiaowei-com-cn/xxwcli.git

git remote add jihulab https://jihulab.com/xuxiaowei-com-cn/xxwcli.git

git remote add github https://github.com/xuxiaowei-com-cn/xxwcli.git

git remote add gitcode https://gitcode.net/xuxiaowei-com-cn/xxwcli.git

git remote add gitlink https://gitlink.org.cn/xuxiaowei-com-cn/xxwcli.git
</details>
