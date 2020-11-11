/*
程序入口
*/

package main

import (
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/server"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
)

func main() {
	if err := chassis.Init(); err != nil {
		openlog.Fatalf("init chassis fail[%s]", err.Error())
	}
	if err := start(); err != nil {
		openlog.Fatalf("system start fail[%s]", err.Error())
	}
	if err := chassis.Run(); err != nil {
		openlog.Fatalf("chassis run fail[%s]", err.Error())
	}
	stop()
}
