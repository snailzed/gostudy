package log

import (
	"testing"
)

var l Log

func TestFileLog(t *testing.T) {
	var err error
	l, err = NewFileLog(LOG_DEBUG, "/Users/apple/go/src/gostudy/demo/runtime/info.log", "/Users/apple/go/src/gostudy/demo/runtime/error.log", SPLIT_BY_SIZE, 1024*1024)
	if err != nil {
		t.Errorf("error:%s", err)
	}
	for {
		l.Debug("debug log")
		l.Info("INfo log")
		l.Warn("Warn log")
		l.Error("Error log")
	}
}
