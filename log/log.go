package log

import (
	l4g "github.com/zdnscloud/cement/log/log4go"
)

type Logger interface {
	Debug(fmt string, args ...interface{})
	Info(fmt string, args ...interface{})
	Warn(fmt string, args ...interface{}) error
	Error(fmt string, args ...interface{}) error
	Close()
}

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
)

func NewLog4jLogger(filename string, level LogLevel, maxSize, maxFileCount int) (Logger, error) {
	return NewLog4jLoggerWithFmt(filename, level, maxSize, maxFileCount, l4g.NewDefaultFormater(l4g.FORMAT_SHORT))
}

func NewLog4jLoggerWithFmt(filename string, level LogLevel, maxSize, maxFileCount int, fmt l4g.LogFormater) (Logger, error) {
	logger := make(l4g.Logger)

	flw, err := l4g.NewFileLogWriter(filename, fmt, maxSize, maxFileCount)
	if err != nil {
		return nil, err
	}

	switch level {
	case Debug:
		logger.AddFilter("file", l4g.DEBUG, flw)
	case Info:
		logger.AddFilter("file", l4g.INFO, flw)
	case Warn:
		logger.AddFilter("file", l4g.WARNING, flw)
	case Error:
		logger.AddFilter("file", l4g.ERROR, flw)
	default:
		panic("unkown level" + string(level))
	}

	return &logger, nil
}

func NewLog4jConsoleLogger(level LogLevel) Logger {
	return NewLog4jConsoleLoggerWithFmt(level, "")
}

func NewLog4jConsoleLoggerWithFmt(level LogLevel, fmt string) Logger {
	if fmt == "" {
		fmt = l4g.FORMAT_SHORT
	}

	switch level {
	case Debug:
		return l4g.NewDefaultLogger(l4g.DEBUG, fmt)
	case Info:
		return l4g.NewDefaultLogger(l4g.INFO, fmt)
	case Warn:
		return l4g.NewDefaultLogger(l4g.WARNING, fmt)
	case Error:
		return l4g.NewDefaultLogger(l4g.ERROR, fmt)
	default:
		panic("unkown level" + string(level))
	}
}

func NewLog4jBufLogger(logChLength uint, level LogLevel) (Logger, chan string) {
	return NewLog4jBufLoggerWithFmt(logChLength, level, l4g.NewDefaultFormater(l4g.FORMAT_SHORT))
}

func NewISO3339Log4jBufLogger(logChLength uint, level LogLevel) (Logger, chan string) {
	formter := &l4g.ISO3339Formator{}
	return NewLog4jBufLoggerWithFmt(logChLength, level, formter)
}

func NewLog4jBufLoggerWithFmt(logChLength uint, level LogLevel, formater l4g.LogFormater) (Logger, chan string) {
	switch level {
	case Debug:
		return l4g.NewBufLogger(logChLength, l4g.DEBUG, formater)
	case Info:
		return l4g.NewBufLogger(logChLength, l4g.INFO, formater)
	case Warn:
		return l4g.NewBufLogger(logChLength, l4g.WARNING, formater)
	case Error:
		return l4g.NewBufLogger(logChLength, l4g.ERROR, formater)
	default:
		panic("unkown level" + string(level))
	}
}

func NewBlackHole() Logger {
	return l4g.NewBlackHole()
}
