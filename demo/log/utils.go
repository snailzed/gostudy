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

func WriteLog(f *FileLog) {
	//for {
	//	select {
	//	case logData := <-f.LogChan:
	//		fmt.Println(logData)
	//		fileStat, err := logData.file.Stat()
	//		if err != nil {
	//			return
	//		}
	//		//判断文件大小是否超出  先重命名，再关闭，然后再创建新的文件句柄
	//		if fileStat.Size() > f.SplitSize {
	//			if logData.error {
	//				reopenFile(f.ErrorPath, f, true)
	//			} else {
	//				reopenFile(f.InfoPath, f, false)
	//			}
	//		}
	//		fmt.Fprintf(logData.file, getFormat(logData.format, logData.level), logData.msg...)
	//	default:
	//		fmt.Println("default")
	//	}
	//}

	for logData := range f.LogChan {
		fmt.Println(logData)
		fileStat, err := logData.file.Stat()
		if err != nil {
			return
		}
		//判断文件大小是否超出  先重命名，再关闭，然后再创建新的文件句柄
		if fileStat.Size() > f.SplitSize {
			if logData.error {
				reopenFile(f.ErrorPath, f, true)
			} else {
				reopenFile(f.InfoPath, f, false)
			}
		}
		fmt.Fprintf(logData.file, getFormat(logData.format, logData.level), logData.msg...)
	}
}

func reopenFile(fpath string, f *FileLog, error bool) {
	if error {
		f.ErrorFile.Close()
	} else {
		f.InfoFile.Close()
	}
	t := time.Now()
	path := fmt.Sprintf("%s_%4d%02d%02d%02d%02d%02d", fpath, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	os.Rename(fpath, path)
	file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	if error {
		f.ErrorFile = file
	} else {
		f.InfoFile = file
	}
}

func getFileLineNum() string {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%s:%s:%d", file, runtime.FuncForPC(pc).Name(), line)
	}
	return ""
}
