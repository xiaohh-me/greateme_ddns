# Golang对接阿里云域名DDNS的项目
## 环境要求
需要 `golang 1.19.*` 的环境，推荐使用 `golang 1.19.9` 版本，各个操作系统的该版本号下载地址：
- Windows amd64：[https://golang.google.cn/dl/go1.19.9.windows-amd64.zip](https://golang.google.cn/dl/go1.19.9.windows-amd64.zip)
- Linux amd64：[https://golang.google.cn/dl/go1.19.9.linux-amd64.tar.gz](https://golang.google.cn/dl/go1.19.9.linux-amd64.tar.gz)
- Linux arm64：[https://golang.google.cn/dl/go1.19.9.linux-arm64.tar.gz](https://golang.google.cn/dl/go1.19.9.linux-arm64.tar.gz)
- MacOS amd64：[https://golang.google.cn/dl/go1.19.9.darwin-amd64.tar.gz](https://golang.google.cn/dl/go1.19.9.darwin-amd64.tar.gz)
- MacOS arm64：[https://golang.google.cn/dl/go1.19.9.darwin-arm64.tar.gz](https://golang.google.cn/dl/go1.19.9.darwin-arm64.tar.gz)
## 设置国内golang代理
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```
## 安装项目依赖
```bash
go mod tidy
```
## 编译代码
注意在Windows下编译的代码只能在Windows下执行，如果需要在Linux下运行则需要在Linux下从新编译，Windows下编译生成的是.exe可执行文件。执行以下代码进行编译（windows下需要安装git并且用git bash执行）：
```bash
./build.sh
```
执行完成后会在项目 `bin` 录地下出现一个可执行文件，文件名各个系统不一样：
- Windows: greateme_ddns.exe
- Linux/MacOS: greateme_ddns

同时会在 `bin/conf` 目录下生成一个 `config.ini` 配置文件，我们需要修改一下这个配置文件：

- `accessKeyId`: 改为阿里云的 `accessKey`
- `accessKeySecret`: 改为阿里云的 `accessKeySecret`
- `domainList`: 域名列表，多个用逗号隔开
- `durationMinute`: 时隔多久更新一次（单位为分钟），默认为十分钟，可无需修改

## 执行代码

可直接在bin目录下执行可执行文件即可