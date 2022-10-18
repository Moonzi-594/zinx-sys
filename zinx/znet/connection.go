package znet

import (
	"fmt"
	"net"
	"project/zinx-sys/zinx/ziface"
)

/*
连接模块
*/

type Connection struct {
	Conn     *net.TCPConn // 连接的socket TCP
	ConnID   uint32       // 连接id
	isClosed bool         // 是否关闭
	ExitChan chan bool    // 告知连接状态的channel

	Router ziface.IRouter // 该连接处理的方法router
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,              // 默认"开启"状态
		ExitChan: make(chan bool, 1), // 有缓冲

		Router: router,
	}
}

// connection的read业务方法
func (c *Connection) StartReader() {
	fmt.Println("[reader] goroutine is running...")
	defer fmt.Printf("[reader] exit! connID=%d, remoteAddr=%s\n", c.ConnID, c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取client数据
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("connection read error:", err)
			continue
		}

		// // 调用connection绑定的HandleAPI
		// if err := c.handleAPI(c.Conn, buf, count); err != nil {
		// 	fmt.Printf("connection handle error:%v, connID=%d", err, c.ConnID)
		// 	break
		// }

		// 得到当前conn的request请求数据
		req := Request{
			conn: c,
			data: buf,
		}

		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		// 从router中找到注册绑定的conn对应的router调用
	}
}

// 启动
func (c *Connection) Start() {
	fmt.Printf("[connection START] connID=%d\n", c.ConnID)

	// 从connection读数据
	go c.StartReader()

	// TODO: write
}

// 停止
func (c *Connection) Stop() {
	fmt.Printf("[connection STOP] connID=%d\n", c.ConnID)

	if c.isClosed {
		return
	}

	c.isClosed = true
	// 关闭socket，释放资源
	c.Conn.Close()
	close(c.ExitChan)
}

// 获取连接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取连接id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取client的 TCP状态 和IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 向client发送数据
func (c *Connection) Send(data []byte) error {
	// TODO:
	return nil
}
