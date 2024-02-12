# shutdown 
![workflow ci](https://github.com/seacraft/shutdown/actions/workflows/ci.yml/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/seacraft/shutdown)](https://pkg.go.dev/github.com/seacraft/shutdown?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/seacraft/shutdown?)](https://goreportcard.com/report/github.com/seacraft/shutdown)
[![golangci badge](https://github.com/golangci/golangci-web/blob/master/src/assets/images/badge_a_plus_flat.svg)](https://golangci.com/r/github.com/seacraft/shutdown)
[![release](https://img.shields.io/github/release-pre/seacraft/shutdown.svg)](https://github.com/seacraft/shutdown/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/seacraft/shutdown/blob/main/LICENSE)
[![Releases](https://img.shields.io/github/downloads/seacraft/shutdown/total.svg)](https://github.com/seacraft/shutdown/releases)

> English | [中文](README_zh.md)

Providing shutdown callbacks for graceful app shutdown

## Installation

```
go get github.com/seacraft/shutdown
```

## Documentation

`github.com/seacraft/shutdown` documentation is available on [godoc](http://godoc.org/github.com/seacraft/shutdown).


## Example - POSIX signals

Graceful shutdown will listen for posix SIGINT and SIGTERM signals. When they are received it will run all callbacks in separate go routines. When callbacks return, the application will exit with os.Exit(0)

```go
package main

import (
	"fmt"
	"time"

	"github.com/seacraft/shutdown"
)

func main() {
	// initialize shutdown
	gs := shutdown.New()

	// add posix shutdown manager
	gs.AddShutdownManager(shutdown.NewPosixSignalManager())

	// add your tasks that implement ShutdownCallback
	gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		fmt.Println("Shutdown callback start")
		time.Sleep(time.Second)
		fmt.Println("Shutdown callback finished")
		return nil
	}))

	// start shutdown managers
	if err := gs.Start(); err != nil {
		fmt.Println("Start:", err)
		return
	}

	// do other stuff
	time.Sleep(time.Hour)
}
```

## Example - posix signals with error handler

The same as above, except now we set an ErrorHandler that prints the error returned from ShutdownCallback.

```go
package main

import (
	"fmt"
	"time"
	"errors"

	"github.com/seacraft/shutdown"
)

func main() {
	// initialize shutdown
	gs := shutdown.New()

	// add posix shutdown manager
	gs.AddShutdownManager(shutdown.NewPosixSignalManager())

	// set error handler
	gs.SetErrorHandler(shutdown.ErrorFunc(func(err error) {
		fmt.Println("Error:", err)
	}))

	// add your tasks that implement ShutdownCallback
	gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		fmt.Println("Shutdown callback start")
		time.Sleep(time.Second)
		fmt.Println("Shutdown callback finished")
		return errors.New("my-error")
	}))

	// start shutdown managers
	if err := gs.Start(); err != nil {
		fmt.Println("Start:", err)
		return
	}

	// do other stuff
	time.Sleep(time.Hour)
}
```
## Licence 

See [LICENSE](LICENSE) file in the root of the repository.
