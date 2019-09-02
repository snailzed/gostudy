package main

import (
	"gostudy/demo/log"
	"gostudy/demo/logger"
	"gostudy/demo/musicSpider"
)

func main1() {
	log := logger.NewFileLogger(logger.LOG_INFO, "./runtime/info.log", "./runtime/warn.log", "./runtime/error.log")
	log.Debug("it is debug log")
	log.Info("it is info log")
	log.Warn("it is warn log")
	log.Error("it is error log")
	log.Fatal("it is fatal log")
}

func main2() {
	var err error
	l, err := log.NewFileLog(log.LOG_DEBUG, "/Users/apple/go/src/gostudy/demo/runtime/info.log", "/Users/apple/go/src/gostudy/demo/runtime/error.log", log.SPLIT_BY_SIZE, 1024*1024)
	if err != nil {
		panic(err)
	}
	for {
		l.Debug("debug log")
	}
}

func main() {
	musicSpider.SpiderJayMusic()
}
