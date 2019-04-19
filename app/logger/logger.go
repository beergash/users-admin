package logger

import (
	"io"
	"log"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger
)

func InitLoggers(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {
	flag := log.Ldate | log.Ltime | log.Lshortfile
	Trace = log.New(traceHandle, "TRACE: ", flag)
	Info = log.New(infoHandle, "INFO: ", flag)
	Warning = log.New(warningHandle, "WARNING: ", flag)
	Error = log.New(errorHandle, "ERROR: ", flag)
}
