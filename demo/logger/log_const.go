package logger

const (
	LOG_DEBUG = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_FATAL
)

var LevelStr map[int]string = map[int]string{
	LOG_DEBUG: "DEBUG",
	LOG_INFO:  "INFO",
	LOG_WARN:  "WARN",
	LOG_ERROR: "ERROR",
	LOG_FATAL: "FATAL",
}
