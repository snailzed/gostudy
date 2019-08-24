package logger

import "testing"

func TestFileLogger(t *testing.T) {
	log := NewFileLogger(LOG_INFO, "/Users/apple/go/src/gostudy/demo/runtime/info.log", "/Users/apple/go/src/gostudy/demo/runtime/warn.log", "/Users/apple/go/src/gostudy/demo/runtime/error.log")
	log.Debug("it is debug log")
	log.Info("it is info log")
	log.Warn("it is warn log")
	log.Error("it is error log")
	log.Fatal("it is fatal log")
}
