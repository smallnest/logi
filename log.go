package logi

import (
	"fmt"
	"log"
	"os"
)

const (
	calldepth = 3
)

var l Logger = &defaultLogger{log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)}

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
}

func SetLogger(logger Logger) {
	l = logger
}

func Debug(v ...interface{}) {
	l.Debug(v)
}
func Debugf(format string, v ...interface{}) {
	l.Debugf(format, v...)
}

func Info(v ...interface{}) {
	l.Info(v)
}
func Infof(format string, v ...interface{}) {
	l.Infof(format, v...)
}

func Warn(v ...interface{}) {
	l.Warn(v)
}
func Warnf(format string, v ...interface{}) {
	l.Warnf(format, v...)
}

func Error(v ...interface{}) {
	l.Error(v)
}
func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	l.Fatal(v)
}
func Fatalf(format string, v ...interface{}) {
	l.Fatalf(format, v...)
}

func Panic(v ...interface{}) {
	l.Panic(v)
}
func Panicf(format string, v ...interface{}) {
	l.Panicf(format, v...)
}

type defaultLogger struct {
	*log.Logger
}

func (l *defaultLogger) Debug(v ...interface{}) {
	l.Output(calldepth, header("[DEBUG]", fmt.Sprint(v...)))
}

func (l *defaultLogger) Debugf(format string, v ...interface{}) {
	l.Output(calldepth, header("[DEBUG]", fmt.Sprintf(format, v...)))
}

func (l *defaultLogger) Info(v ...interface{}) {
	l.Output(calldepth, header("[INFO]", fmt.Sprint(v...)))
}

func (l *defaultLogger) Infof(format string, v ...interface{}) {
	l.Output(calldepth, header("[INFO]", fmt.Sprintf(format, v...)))
}

func (l *defaultLogger) Warn(v ...interface{}) {
	l.Output(calldepth, header("[WARN]", fmt.Sprint(v...)))
}

func (l *defaultLogger) Warnf(format string, v ...interface{}) {
	l.Output(calldepth, header("[WARN]", fmt.Sprintf(format, v...)))
}

func (l *defaultLogger) Error(v ...interface{}) {
	l.Output(calldepth, header("[ERROR]", fmt.Sprint(v...)))
}

func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	l.Output(calldepth, header("[ERROR]", fmt.Sprintf(format, v...)))
}

func (l *defaultLogger) Fatal(v ...interface{}) {
	l.Output(calldepth, header("[FATAL]", fmt.Sprint(v...)))
	os.Exit(1)
}

func (l *defaultLogger) Fatalf(format string, v ...interface{}) {
	l.Output(calldepth, header("[FATAL]", fmt.Sprintf(format, v...)))
	os.Exit(1)
}

func (l *defaultLogger) Panic(v ...interface{}) {
	l.Logger.Panic(v)
}

func (l *defaultLogger) Panicf(format string, v ...interface{}) {
	l.Logger.Panicf(format, v...)
}

func header(lvl, msg string) string {
	return fmt.Sprintf("%s %s", lvl, msg)
}
