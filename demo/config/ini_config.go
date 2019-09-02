package config

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
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
	value := reflect.TypeOf(result)
	if value.Kind() != reflect.Ptr {
		err = errors.New("input a pointer")
		return
	}
	if value.Elem().Kind() != reflect.Struct {
		err = errors.New("input a struct pointer")
		return
	}
	var currentField string
	sectionPattern := regexp.MustCompile(`^\[\s*(.*)\s*\]$`)
	if sectionPattern == nil { //解释失败，返回nil
		fmt.Println("sectionPattern err")
		return
	}
	itemPattern := regexp.MustCompile(`^(.+)\s*=\s*(.*)$`)
	if itemPattern == nil { //解释失败，返回nil
		fmt.Println("sectionPattern err")
		return
	}
	for lineNo, line := range lineArr {
		//判断section是否合法，使用strings包或者正则来匹配
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//注释
		if line[0] == ';' || line[0] == '#' {
			continue
		}
		sectionRet := sectionPattern.FindAllString(line, -1)
		//匹配到了section
		if len(sectionRet) > 0 {
			section := strings.TrimSpace(sectionRet[0])
			section = strings.TrimSpace(section[1 : len(section)-1])
			if len(section) <= 0 {
				err = errors.New(fmt.Sprintf("synax error,section error: %s: %d", line, lineNo+1))
			}
			currentField, err = parseSection(section, value, lineNo)
			continue
		}
		itemRet := itemPattern.FindAllString(line, -1)
		if len(itemRet) == 0 {
			err = errors.New(fmt.Sprintf("synax error, %s: %d", line, lineNo+1))
		}
		//解析Item
		err = parseItem(itemRet[0], currentField, result, lineNo)
		if err != nil {
			return err
		}

	}
	return
}

//解析key=val
func parseItem(line, currentField string, result interface{}, lineNo int) (err error) {

	index := strings.Index(line, "=")
	key := strings.TrimSpace(line[:index])
	value := strings.TrimSpace(line[index+1:])

	fmt.Printf("%s=%s\n", key, value)

	//获取当前的section
	ret := reflect.ValueOf(result)
	section := ret.Elem().FieldByName(currentField)
	if section.Kind() != reflect.Struct {
		err = fmt.Errorf("syanx error: %s must be struct", currentField)
		return
	}
	//遍历section下的字段，判断字段是否存在
	for i := 0; i < section.NumField(); i++ {
		field := section.Type().Field(i)
		itemTag := field.Tag.Get("ini")
		//fmt.Println("item:", itemTag)
		if itemTag == key {
			fmt.Printf("%s.%s=%s\n", currentField, key, value)
			switch section.Field(i).Kind() {
			case reflect.String:
				section.Field(i).SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				section.Field(i).SetInt(valueInt)
			case reflect.Uint:
				valueInt, err := strconv.ParseUint(value, 10, 64)
				if err != nil {
					return err
				}
				section.Field(i).SetUint(valueInt)
			case reflect.Bool:
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				section.Field(i).SetBool(valueBool)
			}

		}
	}
	return
}

func parseSection(parsedSection string, val reflect.Type, lineNo int) (currentField string, err error) {
	for i := 0; i < val.Elem().NumField(); i++ {
		tag := val.Elem().Field(i).Tag.Get("ini")
		if tag == parsedSection {
			currentField = val.Elem().Field(i).Name
			return
		}
	}
	err = errors.New(fmt.Sprintf("Undefined section,%s:%d", parsedSection, lineNo+1))
	return
}
