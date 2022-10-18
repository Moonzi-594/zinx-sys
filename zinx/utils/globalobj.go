package utils

import (
	"encoding/json"
	"io/ioutil"
	"project/zinx-sys/zinx/ziface"
)

/*
存储zinx框架的全局参数，供其他module使用
用户通过zinx.json自行配置
*/

type GlobalObj struct {
	// server
	TcpServer ziface.IServer // zinx的全局Server对象
	Host      string         // 服务器监听的IP地址
	TcpPort   int            // 服务器监听的端口号
	Name      string         // 服务器名称

	// zinx
	Version        string // zinx版本号
	MaxConn        int    // 服务器允许的最大连接数
	MaxPackageSize uint32 // 数据包的容量
}

// 全局的对外Globalobj
var GlobalObject *GlobalObj

func (g *GlobalObj) Load() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// 初始化GlobalObject对象
func init() {
	// 默认值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "beta",
		TcpPort:        8888,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 从.json文件加载用户配置
	GlobalObject.Load()
}
