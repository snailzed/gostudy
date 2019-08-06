package main

import (
	"fmt"
	"net/http"
)

func main() {
	//启动一个http server
	addr := "127.0.0.1:8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("create http server error:", err)
	}
}
