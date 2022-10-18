package ziface

/*
request：连接+数据
*/

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte
}
