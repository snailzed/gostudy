package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"gostudy/demo/protobuf/address"
	"io/ioutil"
)

func main() {
	filename := "./contactBook.txt"
	WriteProto(filename)
	ReadProto(filename)
}

func WriteProto(filename string) {
	var contactBook address.ContactBook
	for i := 0; i < 100; i++ {

		phoneHome := &address.Phone{
			Type:   address.PhoneType_HOME,
			Number: "020-3892987",
		}
		phoneWork := &address.Phone{
			Type:   address.PhoneType_WOKR,
			Number: "+861981234567",
		}
		person := &address.Person{
			Id:     int32(i),
			Name:   fmt.Sprintf("snailZed[%d]", i),
			Phones: []*address.Phone{phoneWork, phoneHome},
		}
		contactBook.Persons = append(contactBook.Persons, person)
	}
	data, err := proto.Marshal(&contactBook)
	if err != nil {
		fmt.Printf("protobuf marshal failed:%v\n", err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("write file failed:%v\n", err)
		return
	}
}

func ReadProto(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file failed:%v\n", err)
		return
	}
	var contactBook address.ContactBook
	err = proto.Unmarshal(content, &contactBook)
	if err != nil {
		fmt.Printf("unmarshal protobuf file failed:%v\n", err)
		return
	}
	fmt.Printf("proto:%+v\n", contactBook)
}
