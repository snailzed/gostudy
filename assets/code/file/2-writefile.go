package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fd, err := os.OpenFile("b.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer fd.Close()
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		_, _ = fd.WriteString(fmt.Sprintf("%d = %d\n", i, i))
	}
	testBufioWriter()
}

func testBufioWriter() {
	fd, err := os.OpenFile("b1.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer fd.Close()
	if err != nil {
		return
	}
	writer := bufio.NewWriter(fd)
	for i := 0; i < 10; i++ {
		_, _ = writer.WriteString(fmt.Sprintf("%d = %d\n", i, i))
	}
	writer.Flush()
}

//例子 请给这个结构编写一个 save 方法，将 Title 作为文件名、Body作为文件内容，写入到文本文件中。
//
//再编写一个 load 函数，接收的参数是字符串 title，该函数读取出与 title 对应的文本文件。请使用 *Page 做为参数，请使用 ioutil 包里的函。
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) saveV1() error {
	fd, err := os.OpenFile(p.Title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	bufferWriter := bufio.NewWriter(fd)
	_, err = bufferWriter.Write(p.Body)
	if err != nil {
		return err
	}
	return nil
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Title, p.Body, 0666)
}

func (p *Page) load(title string) (content []byte, err error) {
	content, err = ioutil.ReadFile(title)
	return
}
