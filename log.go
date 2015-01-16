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

var levels = map[string]Level{}

type Logger struct {
	writer io.Writer
	level  Level
	prefix string
	err    error
	mu     sync.Mutex
}

var Log = &Logger{writer: os.Stdout, level: DEBUG}

func New(w io.Writer, l Level, p string) *Logger {
	return &Logger{writer: w, level: l, prefix: p}
}

func (l *Logger) log(level Level, lvl string, m string, args ...interface{}) {
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

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log(DEBUG, "D", msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log(INFO, "I", msg, args...)
}

func (l *Logger) Warning(msg string, args ...interface{}) {
	l.log(WARNING, "W", msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.log(ERROR, "E", msg, args...)
}
