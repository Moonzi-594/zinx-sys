package main

import (
	"fmt"
	"project/zinx-sys/zinx/ziface"
	"project/zinx-sys/zinx/znet"
)

/*
基于zinx框架的服务端应用程序
*/

// ping tese 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before PING..."))
	if err != nil {
		fmt.Println("call back [before PING] error")
	}
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PING..."))
	if err != nil {
		fmt.Println("call back [PING] error")
	}
}

func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after PING..."))
	if err != nil {
		fmt.Println("call back [after PING] error")
	}
}

func main() {
	s := znet.NewServer("demo-3")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
