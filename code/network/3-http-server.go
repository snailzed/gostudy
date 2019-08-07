package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	startServer()
	//startHttpServer()
}

//
func startServer() {
	//启动一个http server
	addr := "127.0.0.1:8080"
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("create http server error:", err)
	}
}

//v2
func startHttpServer() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(handle))
	if err != nil {
		fmt.Println("create http server error:", err)
	}
}

//处理函数
func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VM", "100")

	//r.Method
	res := `
     r.Method = %s,
     r.URL = %v,
     r.Proto = %v,
     r.ProtoMajor = %v,
     r.ProtoMinor = %v,
`
	fmt.Fprintf(w, res, r.Method, r.URL, r.Proto, r.ProtoMajor, r.ProtoMinor)

	//读取body里面的数据
	body, err := ioutil.ReadAll(r.Body)
	body1, err := ioutil.ReadAll(r.Body)

	fmt.Println(body, body1)
	if err != nil {
		fmt.Println("read body error:", err)
	}
	_, err = io.WriteString(w, string(body))
	_, err = w.Write([]byte(r.URL.String()))
	if err != nil {
		fmt.Println(err)
	}
}
