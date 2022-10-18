package ziface

type IRouter interface {
	// 处理conn业务之前的钩子方法 Hook
	PreHandle(request IRequest)

	// 主业务
	Handle(request IRequest)

	// 处理conn业务之后的钩子方法 Hook
	PostHandle(request IRequest)
}