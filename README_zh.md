# shutdown 
![workflow ci](https://github.com/seacraft/shutdown/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/seacraft/shutdown?)](https://goreportcard.com/report/github.com/seacraft/shutdown)
[![golangci badge](https://github.com/golangci/golangci-web/blob/master/src/assets/images/badge_a_plus_flat.svg)](https://golangci.com/r/github.com/seacraft/shutdown)
[![release](https://img.shields.io/github/release-pre/seacraft/shutdown.svg)](https://github.com/seacraft/shutdown/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/seacraft/shutdown/blob/main/LICENSE)
[![Releases](https://img.shields.io/github/downloads/seacraft/shutdown/total.svg)](https://github.com/seacraft/shutdown/releases)

> [English](README.md) | 中文

提供关闭回调，实现应用程序正常关闭

## Installation

```
go get github.com/seacraft/shutdown
```

## Documentation

`github.com/seacraft/shutdown` documentation 文档可在 [godoc](http://godoc.org/github.com/seacraft/shutdown)上找到。
- [`PosixSignalManager`](http://godoc.org/github.com/seacraft/shutdown/posixsignal)


## 示例 - POSIX 信号

正常关闭将监听 posix SIGINT 和 SIGTERM 信号。当收到它们时，它将在单独的 go 例程中运行所有回调。当回调返回时，应用程序将通过 os.Exit(0) 退出

```go
package main

import (
	"fmt"
	"time"

	"github.com/seacraft/shutdown"
)

func main() {
	// 初始化 shutdown
	gs := shutdown.New()

	// 添加 posix shutdown 管理器
	gs.AddShutdownManager(shutdown.NewPosixSignalManager())

	// 添加实现 ShutdownCallback 的任务
	gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		fmt.Println("Shutdown callback start")
		time.Sleep(time.Second)
		fmt.Println("Shutdown callback finished")
		return nil
	}))

	// 启动shutdown管理器
	if err := gs.Start(); err != nil {
		fmt.Println("Start:", err)
		return
	}

	// 做其他事情
	time.Sleep(time.Hour)
}
```

## 示例 - 带有错误处理程序的 posix 信号

与上面相同，只不过现在我们设置一个 ErrorHandler 来打印从 ShutdownCallback 返回的错误。

```go
package main

import (
	"fmt"
	"time"
	"errors"

	"github.com/seacraft/shutdown"
)

func main() {
	// 初始化 shutdown
	gs := shutdown.New()

	// 添加 posix shutdown 管理器
	gs.AddShutdownManager(shutdown.NewPosixSignalManager())

  // 设置错误处理程序
	gs.SetErrorHandler(shutdown.ErrorFunc(func(err error) {
		fmt.Println("Error:", err)
	}))

	// 添加实现 ShutdownCallback 的任务
	gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		fmt.Println("Shutdown callback start")
		time.Sleep(time.Second)
		fmt.Println("Shutdown callback finished")
		return errors.New("my-error")
	}))

	// 启动shutdown管理器
	if err := gs.Start(); err != nil {
		fmt.Println("Start:", err)
		return
	}

	// 做其他事情
	time.Sleep(time.Hour)
}
```
## 许可证 

请参阅存储库根目录中的 [LICENSE](LICENSE) 
