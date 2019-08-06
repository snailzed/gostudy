package main

import (
	"fmt"
	"net/http"
)

func main() {
	//启动一个http server
	addr := "127.0.0.1:8080"
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("create http server error:", err)
	}
}

//处理函数
func handle(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(r.URL.String()))
	if err != nil {
		fmt.Println(err)
	}
}
