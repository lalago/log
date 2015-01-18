package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

type logger struct {
	writer io.Writer
	level  Level
	prefix string
	err    error
	mu     sync.Mutex
}

var stdLog = &logger{writer: os.Stderr, level: DEBUG}

func New(writer io.Writer, level Level, prefix string) *logger {
	return &logger{writer: writer, level: level, prefix: prefix}
}

func (l *logger) log(level Level, lvl string, m string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > level {
		return
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	log := fmt.Sprintf("%s [%s]%s: %s\n", t, lvl, l.prefix, m)
	if l.err == nil {
		_, err := fmt.Fprintf(l.writer, log, args...)
		if err != nil {
			l.err = err
			fmt.Printf(log, args...)
		}
	}
}

func (l *logger) Debug(msg string, args ...interface{}) {
	l.log(DEBUG, "D", msg, args...)
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.log(INFO, "I", msg, args...)
}

func (l *logger) Warning(msg string, args ...interface{}) {
	l.log(WARNING, "W", msg, args...)
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.log(ERROR, "E", msg, args...)
}

func Debug(msg string, args ...interface{}) {
	stdLog.log(DEBUG, "\033[1;34mD\033[1;m", msg, args...)
}

func Info(msg string, args ...interface{}) {
	stdLog.log(INFO, "\033[1;32mI\033[1;m", msg, args...)
}
func Warning(msg string, args ...interface{}) {
	stdLog.log(WARNING, "\033[1;35mW\033[1;m", msg, args...)
}

func Error(msg string, args ...interface{}) {
	stdLog.log(ERROR, "\033[1;31mE\033[1;m", msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	Error(msg, args...)
	os.Exit(1)
}
