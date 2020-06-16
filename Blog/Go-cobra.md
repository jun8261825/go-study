---
title: Go Cobra
time: 2020年06月16日
auto: 王军
---
# 使用

命令行项目框架生成器。

基本模型

```shell
APPNAME COMMAND ARG --FLAG
```

## 安装

- 下载

```shell
go get -u github.com/spf13/cobra/cobra
```

- 包

```shell
cd $GOPATH/bin
# 可以看到 cobra 包
```

- 初始化

```shell
$GOPATH/bin/cobra init --pkg-name github.com/WalterWj/go-study

# 这样在我的 go-study 项目中，创建了 cmd 目录(初始化会有一个 root.go)和 main.go,LICENSE 两个文件。
```