package main

import (
	"wapi/routes"
	"wapi/server"
)

func main() {
	// 启动rpc服务器
	server.StartRPC()
	// 启动gin服务器
	routes.StartGin()
}
