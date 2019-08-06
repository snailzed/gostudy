package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

//tcp client：读取用户从终端输入的字符，发送到服务器
func main() {
	//readFromTerminalByOsPackage()
	readFromTerminalByBufioPackage()
}

//使用os包读取终端输入
func readFromTerminalByOsPackage() {
	addr := "127.0.0.1:8080"
	//打开连接，创建一个socket
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		////由于目标计算机积极拒绝而无法创建连接
		fmt.Printf("connect to server:%s failed,err=%s", addr, err)
		return
	}
	//读取server返回
	go readFromServer(conn)
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("read error from stdin.")
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("write pack error:", err)
		}
	}
}

func readFromTerminalByBufioPackage() {
	addr := "127.0.0.1:8080"
	//打开连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		////由于目标计算机积极拒绝而无法创建连接
		fmt.Printf("connect to server:%s failed,err=%s", addr, err)
		return
	}
	//读取server返回
	go readFromServer(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, readErr := reader.ReadBytes('\n')
		if readErr != nil {
			fmt.Println("read error from stdin.", readErr)
			return
		}
		input = input[:len(input)-1] //去掉\n ，windows为\r\n
		_, err = conn.Write(input)
		if err != nil {
			fmt.Println("write pack error:", err)
		}
	}
}

//获取从服务端返回的数据
func readFromServer(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("server close.")
				conn.Close()
				os.Exit(-1)
			}
			fmt.Println("read from server error:", err)
			return
		}
		fmt.Println("read from server data:", string(buf[:n]))
	}
}