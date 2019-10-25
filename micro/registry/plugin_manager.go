package registry

import (
	"context"
	"fmt"
	"sync"
)

//初始化
var (
	pluginMgr = &PluginManager{
		plugins: make(map[string]Registry, 2),
	}
)

//注册插件管理类，需要加锁，防止重复初始化
type PluginManager struct {
	plugins map[string]Registry
	lock    sync.Mutex
}

//注册插件
func RegisterPlugin(plugin Registry) (err error) {
	pluginMgr.lock.Lock()
	defer pluginMgr.lock.Unlock()
	if _, ok := pluginMgr.plugins[plugin.Name()]; ok {
		err = fmt.Errorf("duplicate registry plugin")
		return
	}
	pluginMgr.plugins[plugin.Name()] = plugin
	return
}

//初始化插件
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	//查找对应的插件是否存在
	pluginMgr.lock.Lock()
	defer pluginMgr.lock.Unlock()
	plugin, ok := pluginMgr.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exists", name)
		return
	}
	registry = plugin
	err = plugin.Init(ctx, opts...)
	return
}
