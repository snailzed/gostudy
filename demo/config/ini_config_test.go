package config

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Host     string `ini:"host"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Port     int    `ini:"port"`
	Charset  string `ini:"charset"`
	Debug    bool   `ini:"debug"`
}
type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

func TestMarshal(t *testing.T) {
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error(err)
	}
	config := new(Config)
	err = UnMarshal(data, config)
	if err != nil {
		t.Errorf("unmarshal failed , error:%v", err)
	}
	fmt.Printf("%+v\n", config)
}
