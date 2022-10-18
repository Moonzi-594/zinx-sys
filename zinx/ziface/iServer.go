package ziface

// 定义server接口
type IServer interface {
	// 启动
	Start()

	// 停止
	Stop()

	// 运行
	Serve()

	// 路由：为当前的服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(router IRouter)
}
