## 1、环境变量

```bash
setx GOPATH D:\Go
```

这个命令就相当于把`GOPATH`设置到了`D:\Go`下。



### 2、VSCode插件安装

如格式化等，vscode右下角会弹出提示`install`，一版都会下载失败或者下载很慢。

> 在`GOPATH`目录下

按说明配置 go 1.13 及以上，分别执行命令

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

go 1.12及以下，分别执行命令

```shell
# 启用 Go Modules 功能
export GO111MODULE=on
# 配置 GOPROXY 环境变量
export GOPROXY=https://goproxy.io
```

然后从起vscode，点击弹出框的`install`或 `install all`。然后在等待下载完成即可。

```go
Tools environment: GOPATH=D:\Go
Installing 17 tools at D:\Go\bin in module mode.
  gocode
  gopkgs
  go-outline
  go-symbols
  guru
  gorename
  gotests
  gomodifytags
  impl
  fillstruct
  goplay
  godoctor
  dlv
  gocode-gomod
  godef
  goreturns
  golint

Installing github.com/mdempsky/gocode (D:\Go\bin\gocode.exe) SUCCEEDED
Installing github.com/uudashr/gopkgs/v2/cmd/gopkgs (D:\Go\bin\gopkgs.exe) SUCCEEDED
Installing github.com/ramya-rao-a/go-outline (D:\Go\bin\go-outline.exe) SUCCEEDED
Installing github.com/acroca/go-symbols (D:\Go\bin\go-symbols.exe) SUCCEEDED
Installing golang.org/x/tools/cmd/guru (D:\Go\bin\guru.exe) SUCCEEDED
Installing golang.org/x/tools/cmd/gorename (D:\Go\bin\gorename.exe) SUCCEEDED
Installing github.com/cweill/gotests/... (D:\Go\bin\gotests.exe) SUCCEEDED
Installing github.com/fatih/gomodifytags (D:\Go\bin\gomodifytags.exe) SUCCEEDED
Installing github.com/josharian/impl (D:\Go\bin\impl.exe) SUCCEEDED
Installing github.com/davidrjenni/reftools/cmd/fillstruct (D:\Go\bin\fillstruct.exe) SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay (D:\Go\bin\goplay.exe) SUCCEEDED
Installing github.com/godoctor/godoctor (D:\Go\bin\godoctor.exe) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv (D:\Go\bin\dlv.exe) SUCCEEDED
Installing github.com/stamblerre/gocode (D:\Go\bin\gocode-gomod.exe) SUCCEEDED
Installing github.com/rogpeppe/godef (D:\Go\bin\godef.exe) SUCCEEDED
Installing github.com/sqs/goreturns (D:\Go\bin\goreturns.exe) SUCCEEDED
Installing golang.org/x/lint/golint (D:\Go\bin\golint.exe) SUCCEEDED

All tools successfully installed. You are ready to Go :).

```

#### go help modules

报 go: cannot find main module; see 'go help modules'  问题

终端进入项目所在上级目录 执行` go mod init main ` 项目文件夹

会初始化生成一个`go.mod`文件



### 3、常量

常量的注意事项：

- **常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型**
- 不曾使用的常量，在编译的时候，是不会报错的
- 显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，**变量是可以是不同的类型值**