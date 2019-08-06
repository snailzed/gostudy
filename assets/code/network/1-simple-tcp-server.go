package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	fmt.Println("Start to create tcp server...")
	//监听地址
	listerner, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("start server error:", err)
		return
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listerner.Accept()
		if err != nil {
			fmt.Println("Accept connection error:", err)
		}
		go acceptConn(conn)
	}
}

func acceptConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connetion closed.")
				return
			}
			fmt.Println("read data from conn error:", err)
			return
		}
		fmt.Println("read data from conn:", string(buf[:n])) //读取多少字节则输出多少字节
		//发回客户端
		str := strings.ToUpper(string(buf[:n]))
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println("write error:", err)
			_ = conn.Close()
		}
	}
}
