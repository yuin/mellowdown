package log

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"io"
	"sync"
)

type LogLevel int

const (
	Debug LogLevel = iota + 1
	Info
	Warn
	Error
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	default:
		return "ERROR"
	}
}

func (l LogLevel) ColoredString() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "\x1b[36mINFO\x1b[m"
	case Warn:
		return "\x1b[35mWARN\x1b[m"
	default:
		return "\x1b[31mERROR\x1b[m"
	}
}

type Logger interface {
	IsDebugEnabled() bool
	Debug(fmt string, args ...interface{})
	Info(fmt string, args ...interface{})
	Warn(fmt string, args ...interface{})
	Error(fmt string, args ...interface{})
	Write(p []byte) (n int, err error)
}

type logger struct {
	minLevel LogLevel
	stdout   io.Writer
	stderr   io.Writer
	m        sync.Mutex
}

func NewLogger(minLevel LogLevel) Logger {
	return &logger{
		minLevel: minLevel,
		stdout:   colorable.NewColorableStdout(),
		stderr:   colorable.NewColorableStderr(),
	}
}

func (l *logger) Write(p []byte) (int, error) {
	level := Info
	offset := 0
	ln := len(p)
	if ln > 1 && p[0] == '[' {
		switch p[1] {
		case 'D', 'd':
			level = Debug
		case 'I', 'i':
			level = Info
		case 'W', 'w':
			level = Warn
		case 'E', 'e':
			level = Error
		default:
			panic("unknown level:" + string(p))
		}
		for i, b := range p {
			if b == ']' {
				offset = i
				break
			}
		}
	}
	l.p(level, string(p[offset:]))
	return ln, nil
}

func (l *logger) IsDebugEnabled() bool {
	return l.minLevel <= Debug
}

func (l *logger) p(level LogLevel, format string, args ...interface{}) {
	if int(level) < int(l.minLevel) {
		return
	}
	w := l.stdout
	if int(level) >= int(Error) {
		w = l.stderr
	}
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}
	l.m.Lock()
	defer l.m.Unlock()
	fmt.Fprintf(w, "%s\t%s\n", level.ColoredString(), msg)
}

func (l *logger) Debug(format string, args ...interface{}) {
	l.p(Debug, format, args...)
}

func (l *logger) Info(format string, args ...interface{}) {
	l.p(Info, format, args...)
}

func (l *logger) Warn(format string, args ...interface{}) {
	l.p(Warn, format, args...)
}

func (l *logger) Error(format string, args ...interface{}) {
	l.p(Error, format, args...)
}
