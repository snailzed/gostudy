package log

import (
	"os"
)

//异步日志
type FileLog struct {
	InfoPath  string
	ErrorPath string
	Level     int

	InfoFile  *os.File
	ErrorFile *os.File

	LogChan   chan string
	SplitType int
	SplitSize int64
}

func NewFileLog(level int, infoPath, errorPath string, SplitType int, SplitSize int64) (Log, error) {
	f := &FileLog{Level: level, InfoPath: infoPath, ErrorPath: errorPath, SplitType: SplitType, SplitSize: SplitSize}

	fd, err := os.OpenFile(f.InfoPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}
	f.InfoFile = fd
	fd, err = os.OpenFile(f.ErrorPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}
	f.ErrorFile = fd
	return f, nil
}

func (f *FileLog) Debug(format string, args ...interface{}) {
	if f.Level > LOG_DEBUG {
		return
	}
	WriteLog(f.InfoFile, LOG_DEBUG, format, args...)
}
func (f *FileLog) Info(format string, args ...interface{}) {
	if f.Level > LOG_INFO {
		return
	}
	WriteLog(f.InfoFile, LOG_INFO, format, args...)
}
func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.Level > LOG_WARN {
		return
	}
	WriteLog(f.InfoFile, LOG_WARN, format, args...)
}
func (f *FileLog) Error(format string, args ...interface{}) {
	if f.Level > LOG_ERROR {
		return
	}
	WriteLog(f.ErrorFile, LOG_ERROR, format, args...)
}

func (f *FileLog) Close() {
	f.InfoFile.Close()
	f.ErrorFile.Close()
}
