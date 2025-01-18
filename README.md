# Golang对接阿里云域名DDNS的项目
## 环境要求
需要 `golang 1.22.*` 的环境，推荐使用 `golang 1.22.9` 版本，各个操作系统的该版本号下载地址：
- Windows amd64：[https://golang.google.cn/dl/go1.22.9.windows-amd64.zip](https://golang.google.cn/dl/go1.22.9.windows-amd64.zip)
- Linux amd64：[https://golang.google.cn/dl/go1.22.9.linux-amd64.tar.gz](https://golang.google.cn/dl/go1.22.9.linux-amd64.tar.gz)
- Linux arm64：[https://golang.google.cn/dl/go1.22.9.linux-arm64.tar.gz](https://golang.google.cn/dl/go1.22.9.linux-arm64.tar.gz)
- MacOS amd64：[https://golang.google.cn/dl/go1.22.9.darwin-amd64.tar.gz](https://golang.google.cn/dl/go1.22.9.darwin-amd64.tar.gz)
- MacOS arm64：[https://golang.google.cn/dl/go1.22.9.darwin-arm64.tar.gz](https://golang.google.cn/dl/go1.22.9.darwin-arm64.tar.gz)
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
- `domainEndpoint`: 查询域名的Endpoint，默认为杭州，无需修改
- `dnsEndpoint`: DNS的Endpoint，默认为深圳，可根据配置文件当中注释和地理位置进行修改
- `domainList`: 域名列表，多个用逗号隔开
- `dnsType`: 解析类型，只能填写 ipv4 和 ipv6，默认为ipv4（注意全部小写且不能为大写）
- `ipType`: 获取IP地址的类型，可选值：wan 和 interface ，wan：获取当前公网IP地址。interface：根据网卡名称获取IP地址
- `syncWithNoChange`: 在公网IP地址没有改变的时候，是否需要同步，可选值：1 和 0 ，1代表即使IP地址没变化我也需要同步，0则代表在IP地址没变化的情况下不需要同步
- `interfaceName`: 网卡名字，注意填写你设备可用的网卡名称。Windows获取网卡列表命令：ipconfig ，MacOS获取网卡列表命令：ifconfig，Linux获取网卡命令：ip a
- `type`: 执行类型，可选值：single 和 repetition ，single：只执行一次，需要配合系统的定时任务执行。repetition重复执行，需要配合durationMinute配置项执行
- `durationMinute`: 时隔多久更新一次（单位为分钟），默认为十分钟，可无需修改

## 执行代码

可直接在bin目录下执行可执行文件即可