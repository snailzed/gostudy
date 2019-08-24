package logger

import "os"

type ConsoleLogger struct {
	Level int
}

func NewConsoleLogger(level int) Logger {
	return &ConsoleLogger{level}
}

//debug
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.Level > LOG_DEBUG {
		return
	}
	writeLog(os.Stdout, LOG_DEBUG, format, args...)
}

//info
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.Level > LOG_INFO {
		return
	}
	writeLog(os.Stdout, LOG_INFO, format, args...)
}

//warn
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.Level > LOG_WARN {
		return
	}
	writeLog(os.Stdout, LOG_WARN, format, args...)
}

//error
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.Level > LOG_ERROR {
		return
	}
	writeLog(os.Stderr, LOG_ERROR, format, args...)
}

//fatal
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.Level > LOG_FATAL {
		return
	}
	writeLog(os.Stderr, LOG_FATAL, format, args...)
}
