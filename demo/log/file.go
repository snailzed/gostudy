package log

import (
	"fmt"
	"os"
)

//异步日志
type FileLog struct {
	InfoPath  string
	ErrorPath string
	Level     int

	InfoFile  *os.File
	ErrorFile *os.File

	LogChan   chan *LogData
	SplitType int
	SplitSize int64
}

//log data
type LogData struct {
	file   *os.File
	level  int
	format string
	msg    []interface{}
	error  bool
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
	f.LogChan = make(chan *LogData, 50000)
	//异步写入日志
	go WriteLog(f)
	return f, nil
}

func (f *FileLog) Debug(format string, args ...interface{}) {
	if f.Level > LOG_DEBUG {
		return
	}
	logData := &LogData{
		file:   f.InfoFile,
		level:  LOG_DEBUG,
		format: format,
		msg:    args,
		error:  false,
	}
	//select通道选择器：当通道阻塞时，会走default语句，不会出现阻塞现象，这时会丢弃日志
	select {
	case f.LogChan <- logData:
		fmt.Println("IN")
	default:
		fmt.Println("DEFAULT")
	}
}
func (f *FileLog) Info(format string, args ...interface{}) {
	if f.Level > LOG_INFO {
		return
	}
	logData := &LogData{
		file:   f.InfoFile,
		level:  LOG_INFO,
		format: format,
		msg:    args,
		error:  false,
	}
	//select通道选择器：当通道阻塞时，会走default语句，不会出现阻塞现象，这时会丢弃日志
	select {
	case f.LogChan <- logData:
	default:
	}
}
func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.Level > LOG_WARN {
		return
	}
	logData := &LogData{
		file:   f.InfoFile,
		level:  LOG_WARN,
		format: format,
		msg:    args,
		error:  false,
	}
	//select通道选择器：当通道阻塞时，会走default语句，不会出现阻塞现象，这时会丢弃日志
	select {
	case f.LogChan <- logData:
	default:
	}
}
func (f *FileLog) Error(format string, args ...interface{}) {
	if f.Level > LOG_ERROR {
		return
	}
	logData := &LogData{
		file:   f.ErrorFile,
		level:  LOG_ERROR,
		format: format,
		msg:    args,
		error:  true,
	}
	//select通道选择器：当通道阻塞时，会走default语句，不会出现阻塞现象，这时会丢弃日志
	select {
	case f.LogChan <- logData:
	default:
	}
}

func (f *FileLog) Close() {
	f.InfoFile.Close()
	f.ErrorFile.Close()
}
