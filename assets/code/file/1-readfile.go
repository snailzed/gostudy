package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//readFileBybufio()
	//readFileByioutil()
	//readFileBybufioBuffer()
	//readFileByFmt()
	readCsv("./a.txt")
}

//使用bufio循环读取文件
func readFileBybufio() {
	//1、打开一个文件
	fd, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer fd.Close()

	//2、使用bufio读取文件
	reader := bufio.NewReader(fd)
	for {
		inputString, readerError := reader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

//使用ioutil一次性读取文件
func readFileByioutil() {
	//1、使用ioutil一次性读取string
	content, err := ioutil.ReadFile("./a.txt")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	fmt.Println(string(content))
}

func readFileBybufioBuffer() {
	fd, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fd.Close()

	//带缓冲循环读取文件内容
	//var buffer []byte; //切片
	buffer := make([]byte, 1024)
	for {
		n, err := fd.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("read file error:", err)
			return
		}
		if n == 0 {
			break
		}
	}
	fmt.Println(string(buffer))
}

func readFileByFmt() {
	fd, err := os.Open("./a.txt")
	if err != nil {
		return
	}
	var line string
	fmt.Fscan(fd, &line)
	fmt.Println(line)
}

func readCsv(file string) {
	fd, err := os.Open(file)
	if err != nil {
		return
	}
	reader := bufio.NewReader(fd)
	data := make([]map[string]string, 4)
	i := 0
	for {
		line, readErr := reader.ReadString('\n')
		line = line[:len(line)-1]
		d := strings.Split(line, ";")
		pro := make(map[string]string, 3)
		pro["title"] = d[0]
		pro["price"] = d[1]
		pro["quantity"] = d[2]

		data[i] = pro
		i++
		if readErr == io.EOF {
			break
		}
	}
	fmt.Println(data)
}
