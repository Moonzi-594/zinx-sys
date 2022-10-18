package znet

import (
	"fmt"
	"net"
	"project/zinx-sys/zinx/ziface"
)

// IServer的接口实现
type Server struct {
	Name      string
	IPVersion string // tcp, udp...
	IP        string
	Port      int
	Router    ziface.IRouter
}

// 定理client连接绑定的handle API TODO: 未来会优化，由用户自定义
// func CallackToClient(conn *net.TCPConn, data []byte, count int) error {
// 	// 回显业务
// 	fmt.Println("connection handle callback!")
// 	if _, err := conn.Write(data[:count]); err != nil {
// 		fmt.Println("CallbackToClient Error:", err)
// 		return errors.New("CallbackToClient Error")
// 	}

// 	return nil
// }

func (s *Server) Start() {
	fmt.Printf("[START] The server is listening on port %d at address %s...\n", s.Port, s.IP)

	go func() {
		// 1.获取TCPAddr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve TCP addr error:", err)
			return
		}

		// 2.监听server addr
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listening ", s.IPVersion, " error ", err)
			return
		}
		fmt.Printf("Successfully started zinx service [%s], listening...\n", s.Name)

		var cid uint32 = 0

		// 3.阻塞等待client连接，处理client业务（read & write）
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error:", err)
				return
			}

			// 将业务方法与conn绑定，得到connection模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 启动连接业务
			go dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {
	// TODO: 停止server资源、状态、连接信息等
}

func (s *Server) Serve() {
	s.Start() // 异步

	// TODO: 额外业务

	// 阻塞server
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Success!")
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
		Router:    nil,
	}
}
