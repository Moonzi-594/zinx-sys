package main

import "project/zinx-sys/zinx/znet"

/*
基于zinx框架的服务端应用程序
*/

func main() {
	s := znet.NewServer("demo-2")
	s.Serve()
}
