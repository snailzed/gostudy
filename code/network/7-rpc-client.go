package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Cal struct {
	N1 int
	N2 int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	//同步调用
	var reply int
	cal := &Cal{5, 6}
	err = client.Call("Cal.Multi", cal, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d\n", cal.N1, cal.N2, reply)
}
