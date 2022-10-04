package ziface

import "net"

type IConnection interface {
	// 启动
	Start()

	// 停止
	Stop()

	// 获取连接绑定的socket conn
	GetTCPConnection() *net.TCPConn

	// 获取连接id
	GetConnID() uint32

	// 获取client的 TCP状态 和IP port
	RemoteAddr() net.Addr

	// 向client发送数据
	Send(data []byte) error
}

// 处理连接对应的业务
type HandleFunc func(*net.TCPConn, []byte, int) error
