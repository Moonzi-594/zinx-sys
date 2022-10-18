package znet

import "project/zinx-sys/zinx/ziface"

// 实现router时，先嵌入该BaseRouter基类，再根据需要重写
type BaseRouter struct{}

// 处理conn业务之前的钩子方法 Hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

// 主业务
func (br *BaseRouter) Handle(request ziface.IRequest) {}

// 处理conn业务之后的钩子方法 Hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
