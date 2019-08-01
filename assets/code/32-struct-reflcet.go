package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string "it is name"
	age  int    "it is age"
	desc string "it is desc"
}

func main() {
	a := new(Animal)
	atype := reflect.TypeOf(*a)
	fmt.Println("field 0 type = ", atype.Field(0).Tag)
	fmt.Println("field 1 type = ", atype.Field(1).Tag)
	fmt.Println("field 2 type = ", atype.Field(2).Tag)
	/**
	field 0 type =  it is name
	field 1 type =  it is age
	field 2 type =  it is desc
	*/
}
