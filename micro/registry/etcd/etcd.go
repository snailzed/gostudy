package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"gostudy/micro/registry"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxServiceNum          = 8
	MaxSyncServiceInterval = 10
)

//etcd注册中心
type EtcdRegistry struct {
	options   *registry.Options
	client    *clientv3.Client
	serviceCh chan *registry.Service

	value              atomic.Value
	lock               sync.Mutex
	registryServiceMap map[string]*RegisterService
}

type RegisterService struct {
	id          clientv3.LeaseID
	service     *registry.Service
	registered  bool
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
}

type AllServiceInfo struct {
	serviceMap map[string]*registry.Service
}

var (
	etcdRegistry = &EtcdRegistry{
		serviceCh:          make(chan *registry.Service, MaxServiceNum),
		registryServiceMap: make(map[string]*RegisterService, MaxServiceNum),
	}
)

func init() {

	allServiceInfo := &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaxServiceNum),
	}

	etcdRegistry.value.Store(allServiceInfo)
	registry.RegisterPlugin(etcdRegistry)
	go etcdRegistry.run()
}

//插件名字
func (etcd *EtcdRegistry) Name() string {
	return "etcd"
}

//初始化
func (etcd *EtcdRegistry) Init(ctx context.Context, opts ...registry.Option) (err error) {
	etcd.options = &registry.Options{}
	for _, f := range opts {
		f(etcd.options)
	}
	//连接etcd
	etcd.client, err = clientv3.New(clientv3.Config{
		Endpoints:   etcd.options.Addrs,
		DialTimeout: etcd.options.Timeout,
	})
	if err != nil {
		// handle error!
		err = fmt.Errorf("init etcd failed, err:%v", err)
		return
	}
	return
}

//注册服务
func (etcd *EtcdRegistry) Register(ctx context.Context, service *registry.Service) (err error) {
	return
}

//反注册
func (etcd *EtcdRegistry) Unregister(ctx context.Context, service *registry.Service) (err error) {
	return
}

func (e *EtcdRegistry) run() {

	ticker := time.NewTicker(MaxSyncServiceInterval)
	for {
		select {
		case service := <-e.serviceCh:
			registryService, ok := e.registryServiceMap[service.Name]
			if ok {
				for _, node := range service.Nodes {
					registryService.service.Nodes = append(registryService.service.Nodes, node)
				}
				registryService.registered = false
				break
			}
			registryService = &RegisterService{
				service: service,
			}
			e.registryServiceMap[service.Name] = registryService
		case <-ticker.C:
			//e.syncServiceFromEtcd()
		default:
			//e.registerOrKeepAlive()
			time.Sleep(time.Millisecond * 500)
		}
	}
}
