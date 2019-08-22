package log

import "testing"

func TestFileLog(t *testing.T) {
	l, err := NewFileLog(LOG_DEBUG, "/Users/snailzed/go/src/gostudy/demo/runtime/info.log", "/Users/snailzed/go/src/gostudy/demo/runtime/error.log")
	if err != nil {
		t.Errorf("error:%s", err)
	}
	l.Debug("debug log")
	l.Info("INfo log")
	l.Warn("Warn log")
	l.Error("Error log")
}
