package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func Format(level int) string {
	return fmt.Sprintf("[%s] [%s] ", time.Now().Format("2006-01-02 15:04:05"), LevelStr[level])
}

func getFileLineNum() string {
	//实参skip为上溯的栈帧数，0表示Caller的调用者（Caller所在的调用栈
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf(" call %s in file %s at line %d", runtime.FuncForPC(pc).Name(), file, line)
	}
	return ""
}

func writeLog(file *os.File, level int, format string, args ...interface{}) {
	format = Format(level) + format + getFileLineNum() + "\n"
	fmt.Fprintf(file, format, args...)
}
