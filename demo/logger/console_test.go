package logger

import (
	"testing"
)

func TestConsole(t *testing.T) {
	log := NewConsoleLogger(LOG_DEBUG)
	log.Debug("it is debug log")
	log.Info("it is info log")
	log.Warn("it is warn log")
	log.Error("it is error log")
	log.Fatal("it is fatal log")
}
