package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

const (
	autoColor = iota
	disableColor
	forceColor
)

var (
	ColorGreen       = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	ColorWhite       = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	ColorYellow      = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	ColorRed         = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	ColorBlue        = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	ColorMagenta     = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	ColorCyan        = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	ColorReset       = string([]byte{27, 91, 48, 109})
	consoleColorMode = autoColor
)

var gLogTag map[int]string = map[int]string{}

func init() {
	gLogTag[DEBUG] = "[Debug]"
	gLogTag[INFO] = "[Info]"
	gLogTag[WARN] = "[Warn]"
	gLogTag[ERROR] = "[Error]"
	gLogTag[FATAL] = "[Fatal]"
}

type LogFunc func(level int, format string, args ...interface{})

var Log LogFunc = DefaultLogf

// Level default log level is DEBUG
var Level = DEBUG

var Writer io.Writer = os.Stdout

// InnerLevel inner log level is DEBUG
var InnerLevel = DEBUG

func DefaultLogf(level int, format string, args ...interface{}) {
	if Level > level {
		return
	}

	logInfo := ""
	length := len(format)
	if length > 0 {
		if format[length-1] != '\n' {
			format = format + "\n"
		}
		logInfo = fmt.Sprintf(format, args...)
	} else {
		logInfo = fmt.Sprintln(args...)
	}

	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	Writer.Write([]byte(fmt.Sprintf("%-7s %s %s:%d %s", gLogTag[level], TimeFormat(time.Now()), shortFile(file), line, logInfo)))
	if level >= FATAL {
		os.Exit(-1)
	}
}

func shortFile(file string) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short
}

// Debug log level.
//  @param format log format
//  @param args   log args
//  usage: log.Debug("test")             => [Debug]  2022-06-22 - 13:14:23 log_test.go:11 test
//  usage: log.Debug("%d %s", 5, "test") => [Debug]  2022-06-22 - 13:14:23 log_test.go:11 5 test
func Debug(format string, args ...interface{}) {
	Log(DEBUG, format, args...)
}

// Info log level.
//  @param format log format
//  @param args   log args
//  usage: log.Info("test")             => [Info]  2022-06-22 - 13:14:23 log_test.go:11 test
//  usage: log.Info("%d %s", 5, "test") => [Info]  2022-06-22 - 13:14:23 log_test.go:11 5 test
func Info(format string, args ...interface{}) {
	Log(INFO, format, args...)
}

// Warn log level.
//  @param format log format
//  @param args   log args
//  usage: log.Warn("test")             => [Warn]  2022-06-22 - 13:14:23 log_test.go:11 test
//  usage: log.Warn("%d %s", 5, "test") => [Warn]  2022-06-22 - 13:14:23 log_test.go:11 5 test
func Warn(format string, args ...interface{}) {
	Log(WARN, format, args...)
}

// Error log level.
//  @param format log format
//  @param args   log args
//  usage: log.Error("test")             => [Error]  2022-06-22 - 13:14:23 log_test.go:11 test
//  usage: log.Error("%d %s", 5, "test") => [Error]  2022-06-22 - 13:14:23 log_test.go:11 5 test
func Error(format string, args ...interface{}) {
	Log(ERROR, format, args...)
}

// Fatal log level.
//  @param format log format
//  @param args   log args
//  usage: log.Fatal("test")             => [Fatal]  2022-06-22 - 13:14:23 log_test.go:11 test
//  usage: log.Fatal("%d %s", 5, "test") => [Fatal]  2022-06-22 - 13:14:23 log_test.go:11 5 test
func Fatal(format string, args ...interface{}) {
	Log(FATAL, format, args...)
}

// InnerLog log level, the level dependence on your log level.
//  @param format log format
//  @param args   log args
//  usage: log.InnerLog("test")             => [Debug]  2022-06-22 - 13:14:23 log_test.go:11 test
//  usage: log.InnerLog("%d %s", 5, "test") => [Debug]  2022-06-22 - 13:14:23 log_test.go:11 5 test
func InnerLog(format string, args ...interface{}) {
	Log(InnerLevel, format, args...)
}

func TimeFormat(t time.Time) string {
	var timeString = t.Format("2006-01-02 - 15:04:05")
	return timeString
}
