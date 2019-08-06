package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

//
type A struct {
	X, y, z int
	Name    string
}

//
type B struct {
	X, Y int
	Name string
	int
}

//
func main() {
	writeToFile()
	readFromFile()
	return
	var network bytes.Buffer

	//编码
	enc := gob.NewEncoder(&network)
	_ = enc.Encode(A{1, 2, 3, ""})

	//解码
	var b B
	dec := gob.NewDecoder(&network)
	_ = dec.Decode(&b)
	fmt.Println(b)
}

func writeToFile() {
	b := B{1, 2, "Snailzed", 1}
	fd, err := os.OpenFile("test.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer fd.Close()
	if err != nil {
		return
	}
	encode := gob.NewEncoder(fd)
	err = encode.Encode(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readFromFile() {
	fd, err := os.Open("test.gob")
	defer fd.Close()
	if err != nil {
		return
	}
	dec := gob.NewDecoder(fd)
	var a A
	err = dec.Decode(&a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}
