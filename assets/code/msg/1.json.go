package main

import (
	"encoding/json"
	"fmt"
)

type Msg struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	msg := Msg{Code: 0, Message: "Hello", Data: 1}
	m, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m, string(m))

	var map1 map[string]interface{}
	err = json.Unmarshal(m, &map1)
	fmt.Println(map1)

	var m2 Msg
	err = json.Unmarshal(m, &m2)
	fmt.Println(m2)
}
