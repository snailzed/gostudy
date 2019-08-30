package logger

import (
	"os"
)

//文件日志结构体
type FileLogger struct {
	Level         int
	InfoFilePath  string
	WarnFilePath  string
	ErrorFilePath string

	InfoFile  *os.File
	WarnFile  *os.File
	ErrorFile *os.File
}

//new instance
func NewFileLogger(level int, InfoFilePath, WarnFilePath, ErrorFilePath string) Logger {
	f := &FileLogger{Level: level}
	fd, err := os.OpenFile(InfoFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("open file error. path=" + InfoFilePath)
	}
	f.InfoFile = fd
	fd, err = os.OpenFile(WarnFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("open file error. path=" + WarnFilePath)
	}
	f.WarnFile = fd
	fd, err = os.OpenFile(ErrorFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("open file error. path=" + ErrorFilePath)
	}
	f.ErrorFile = fd
	return f
}

//debug
func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.Level > LOG_DEBUG {
		return
	}
	writeLog(f.InfoFile, LOG_DEBUG, format, args...)
}

//info
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.Level > LOG_INFO {
		return
	}
	writeLog(f.InfoFile, LOG_INFO, format, args...)
}

//warn
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.Level > LOG_WARN {
		return
	}
	writeLog(f.WarnFile, LOG_WARN, format, args...)
}

//error
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.Level > LOG_ERROR {
		return
	}
	writeLog(f.ErrorFile, LOG_ERROR, format, args...)
}

//fatal
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.Level > LOG_FATAL {
		return
	}
	writeLog(f.ErrorFile, LOG_FATAL, format, args...)
}

//close
func (f *FileLogger) Close() {
	_ = f.InfoFile.Close()
	_ = f.WarnFile.Close()
	_ = f.ErrorFile.Close()
}
