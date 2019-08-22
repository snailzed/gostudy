package log

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func getFormat(format string, level int) string {
	return fmt.Sprintf("[%s] [%s] [%s] ", time.Now().Format("2006-01-02 15:04:05"), LogStr[level], getFileLineNum()) + format + "\n"
}

func WriteLog(file *os.File, level int, format string, args ...interface{}) {
	fmt.Fprintf(file, getFormat(format, level), args...)
}

func getFileLineNum() string {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%s:%s:%d", file, runtime.FuncForPC(pc).Name(), line)
	}
	return ""
}
