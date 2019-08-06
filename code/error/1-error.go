package main

import (
	"errors"
	"fmt"
)

func main() {

	//1、errors.New() 返回 *errors.errorString对象
	err := errors.New("Error")
	fmt.Println(err)
	fmt.Printf("%t,%T\n", err, err)

	//2、fmt.Errorf 返回 *errors.errorString对象
	err1 := fmt.Errorf("%s%d", "Error:", 1)
	fmt.Println(err1)
	fmt.Printf("%t,%T", err1, err1)

}
