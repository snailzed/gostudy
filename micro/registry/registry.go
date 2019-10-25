package registry

import "context"

//服务注册插件接口，ctx用来传递上下文
type Registry interface {
	//插件名字
	Name() string
	//初始化
	Init(ctx context.Context, opts ...Option) (err error)
	//注册服务
	Register(ctx context.Context, service *Service) (err error)
	//反注册
	Unregister(ctx context.Context, service *Service) (err error)
}
