package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

type ConfigXml struct {
	Name       xml.Name        `xml:"config"`
	Version    string          `xml:"version,attr"`
	ServerConf ServerConfigXml `xml:"server"`
	MysqlConf  MysqlConfigXml  `xml:"mysql"`
}

type ServerConfigXml struct {
	Name string `xml:"name"`
	Ip   string `xml:"ip"`
	Port int    `xml:"port"`
}

type MysqlConfigXml struct {
	Host     string `xml:"host"`
	Port     int    `xml:"port"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Database string `xml:"database"`
}

func TestUnMarshalXml(t *testing.T) {
	//读取配置文件内容
	data, err := ioutil.ReadFile("./config.xml")
	if err != nil {
		fmt.Println("read file error:", err)
	}
	configXml := new(ConfigXml)
	//解析
	err = xml.Unmarshal(data, configXml)
	if err != nil {
		fmt.Printf("Parse xml file error:%v\n", err)
		return
	}
	fmt.Printf("xml:%+v\n", configXml)

	xmlStr, err := xml.Marshal(configXml)
	if err != nil {
		fmt.Printf("error:%s\n", err)
	}
	fmt.Printf(string(xmlStr))
}
