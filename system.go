/*
 系统启动、停止执行
*/
package main

import (
	"fmt"

	"github.com/go-chassis/go-archaius"
	_ "github.com/go-chassis/go-archaius/source/remote/kie"
	"github.com/go-chassis/openlog"

	_ "github.com/wksw/go-chassis-demo/handlers"
)

// 启动
func start() error {
	openlog.Info("system starting...")
	// TODO something
	// 检查是否连接到了配置中心
	if archaius.GetString("project", "") == "" {
		openlog.Fatal(fmt.Sprintf("connect to config center fail"))
	}
	openlog.Info("system started")
	return nil
}

// 停止
func stop() {
	openlog.Info("system stoping...")
	// TODO something
	openlog.Info("system stoped, bye bye!")
}
