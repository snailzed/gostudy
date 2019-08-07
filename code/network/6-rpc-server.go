package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

//定义一个cal服务  func (t *T) MethodName(argType T1, replyType *T2) error
type Cal struct {
	N1 int
	N2 int
}

//乘法
func (c *Cal) Multi(ca *Cal, reply *int) error {
	*reply = ca.N1 * ca.N2
	return nil
}

//除法
func (c *Cal) Div(ca *Cal, reply *int) error {
	if ca.N2 == 0 {
		return errors.New("error")
	}
	*reply = ca.N1 / ca.N2
	return nil
}

func main() {
	cal := new(Cal)          //必须是指针
	err := rpc.Register(cal) //注册服务
	//rpc.RegisterName("cal",cal)//使用名称注册服务
	if err != nil {
		fmt.Print("register service error:", err)
		return
	}
	rpc.HandleHTTP() //处理http请求
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Print("listen tcp server error:", err)
		return
	}
	go http.Serve(listen, nil) //会为每一个进入 listener 的 HTTP 连接创建新的服务线程
	time.Sleep(time.Second * 30)
}
