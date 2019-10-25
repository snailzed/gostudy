package registry

import (
	"time"
)

//选项模式: 返回匿名函数,通过执行一系列的匿名函数来初始化选项
type Options struct {
	Addrs   []string
	Timeout time.Duration
	// example:  /xxx_company/app/kuaishou/service_A/10.192.1.1:8801
	// example:  /xxx_company/app/kuaishou/service_A/10.192.1.2:8801
	RegistryPath string
	HeartBeat    int64
}

//选项函数
type Option func(opts *Options)

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func WithRegistryPath(path string) Option {
	return func(opts *Options) {
		opts.RegistryPath = path
	}
}

func WithHeartBeat(heartHeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartHeat
	}
}

//初始化
//
//func init(opfuc option...) {
//   opts = new(Options)
//   for f := range opfuc{
//      f(opts)
//   }
// }
//
//
//
//
