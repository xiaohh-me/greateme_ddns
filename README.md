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
## 修改项目代码
我们需要修改项目根目录当中 `main.go` 的代码，具体在常量定义(文件第12~21行)当中的几个变量：
- AccessKeyId: 阿里云的AccessKey，需要在阿里云当中申请
- AccessKeySecret: 阿里云的AccessKeySecret，需要在阿里云当中申请
- SyncDomainList: 需要同步的全域名，多个用英文逗号","隔开
- DurationMinute: 多久同步一次，单位为分钟，代码中默认写着的是十分钟
## 编译代码
注意在Windows下编译的代码只能在Windows下执行，如果需要在Linux下运行则需要在Linux下从新编译，Windows下编译生成的是.exe可执行文件。执行以下代码进行编译：
```bash
go build .
```
执行完成后会在项目根目录地下出现一个可执行文件，文件名各个系统不一样：
- Windows: greateme_ddns.exe
- Linux/MacOS: greateme_ddns
生成完成后直接执行这个可执行文件即可进行域名同步