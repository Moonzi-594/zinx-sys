package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
在linux中使用nc也可以模拟客户端连接
netcat: a utility for sending raw data over a network connection.
=> Arbitrary TCP and UDP connections and listens
*/

func main() {
	fmt.Println("[START] The client is start...")
	time.Sleep(1 * time.Second)

	// 1.连接server，得到connection
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}

	for {
		// 2.写数据
		var clientMsg string
		fmt.Println("请输入发送至server的数据:")
		fmt.Scanln(&clientMsg)

		_, err := conn.Write([]byte(clientMsg))
		if err != nil {
			fmt.Println("conn.Write error:", err)
			return
		}

		buf := make([]byte, 4096)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read error:", err)
			return
		}
		fmt.Printf("[echo]%s\n", buf)

		time.Sleep(1 * time.Second)
	}

}
