package log

import (
	"fmt"
	"log"
	"os"
)

const (
	FATAL = iota
	PANIC
	ERROR
	WARN
	INFO
	DEBUG
)

type defaultLogger struct {
	*log.Logger
	Level int
}

func (l *defaultLogger) output(level int, levelKey string, v ...interface{}) error {
	if l.Level >= level {
		h := header(levelKey, fmt.Sprint(v...))
		return l.Output(calldepth, h)
	}else {
		return nil
	}
}

func (l *defaultLogger) outputf(level int, levelKey string, format string, v ...interface{}) error {
	if l.Level >= level {
		h := header(levelKey, fmt.Sprintf(format, v...))
		return l.Output(calldepth, h)
	} else {
		return nil
	}
}


func (l *defaultLogger) Debug(v ...interface{}) {
	l.output(DEBUG, "DEBUG", v)
}

func (l *defaultLogger) Debugf(format string, v ...interface{}) {
	l.outputf(DEBUG, "DEBUG", format, v...)
}

func (l *defaultLogger) Info(v ...interface{}) {
	l.output(INFO, "INFO", v)
}

func (l *defaultLogger) Infof(format string, v ...interface{}) {
	l.outputf(INFO, "INFO", format, v...)
}

func (l *defaultLogger) Warn(v ...interface{}) {
	l.output(WARN, "WARN", v)
}

func (l *defaultLogger) Warnf(format string, v ...interface{}) {
	l.outputf(WARN, "WARN", format, v...)
}

func (l *defaultLogger) Error(v ...interface{}) {
	l.output(ERROR, "ERROR", v)
}

func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	l.outputf(ERROR, "ERROR", format, v...)
}

func (l *defaultLogger) Fatal(v ...interface{}) {
	l.output(FATAL, "FATAL", v)
	os.Exit(1)
}

func (l *defaultLogger) Fatalf(format string, v ...interface{}) {
	l.outputf(FATAL, "FATAL", format, v...)
	os.Exit(1)
}

func (l *defaultLogger) Panic(v ...interface{}) {
	l.Logger.Panic(v)
}

func (l *defaultLogger) Panicf(format string, v ...interface{}) {
	l.Logger.Panicf(format, v...)
}

func header(lvl, msg string) string {
	return fmt.Sprintf("%s: %s", lvl, msg)
}
