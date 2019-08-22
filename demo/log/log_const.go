package log

const (
	LOG_DEBUG = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
)

const (
	SPLIT_BY_SIZE = iota
	SPLIT_BY_HOUR
)

var LogStr map[int]string = map[int]string{
	LOG_DEBUG: "DEBUG",
	LOG_INFO:  "INFO",
	LOG_WARN:  "WARN",
	LOG_ERROR: "ERROR",
}
