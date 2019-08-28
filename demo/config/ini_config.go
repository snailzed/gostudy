package config

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

//将结构体类型转换成 ini格式字节数组，返回字节数组
func Marshal(data interface{}) (result []byte, err error) {
	return
}

//将ini类型字节数组转换成结构体
// 读取ini配置文件内容--》解析文件内容 --》设置到对应结构体属性值
//1、获取结构体所有字段，遍历
//2、获取每个字段对应的tag，判断当前行是否一个配置选项
//3、解析每个大配置项下面的选项,并设置结构体值

func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	value := reflect.ValueOf(result)
	if value.Type().Kind() != reflect.Ptr {
		err = errors.New("input a pointer")
		return
	}
	if value.Elem().Kind() != reflect.Struct {
		err = errors.New("input a struct pointer")
		return
	}
	for _, line := range lineArr {
		fmt.Println(line)
		//判断section是否合法，使用strings包或者正则来匹配
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		for i := 0; i < value.NumField(); i++ {

		}
	}
	return
}
