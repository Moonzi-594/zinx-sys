package znet

import (
	"fmt"
	"net"
	"project/zinx-sys/zinx/ziface"
)

// IServer的接口实现
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

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

		// 3.阻塞等待client连接，处理client业务（read & write）
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error:", err)
				return
			}

			// client-server连接建立，处理业务
			go func() { // 回显：max-512KB
				for {
					buf := make([]byte, 512)
					count, err := conn.Read(buf)
					if err != nil {
						fmt.Println("receive buf error:", err)
						continue
					}

					// server端显示收到的信息
					fmt.Printf("[FROM client]%s\n", buf)

					if _, err := conn.Write(buf[:count]); err != nil {
						fmt.Println("write back error:", err)
						continue
					}
				}
			}()
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

func NewServer(name string) ziface.IServer {
	server := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
	}

	return server
}
