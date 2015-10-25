# log6

Log6 is a simple wrapper of the go log library.

Log6 provide Debug, Info, Warn, Error, Fatal, Panic functions to log information.

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your `test.go` file.

~~~ go
package main

import (
    "os"
    "github.com/rod6/log6"
)

func main() {
    f, _ := os.OpenFile("testing.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
    defer f.Close()

    log6.SetOutput(f)
    log6.SetLevel(log6.InfoLevel)

    log6.Debug("Log Debug Message: %v", "debug")
    log6.Info("Log Info Message: %v", "info")
    log6.Warn("Log Warn Message: %v", "warn")
    log6.Error("Log Error Message: %v", "error")
}
~~~

Then install the Log6 package (**go 1.1** and greater is required):
~~~
go get github.com/rod6/log6
~~~

Then run your test:
~~~
go run test.go
~~~

Then you can check the log file: testing.log.
