package registry

//服务定义
type Service struct {
	Name  string  //服务名
	Nodes []*Node //服务节点
}

//节点定义
type Node struct {
	Id   string //id
	Ip   string //ip地址
	Port int    //ip端口
}
