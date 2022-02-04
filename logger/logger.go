package logger

import (
	"io"
	"log"
	"sync"
)

var (
	Nil = Logger{}

	once     sync.Once
	instance Logger
)

type ILogger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
}

type Logger struct {
	logger *log.Logger
}

// New returns a new Logger.
func New(w io.Writer) *Logger {
	return &Logger{
		logger: log.New(w, "", log.Ldate|log.Ltime|log.Lmsgprefix),
	}
}

// NewSingleton returns an Instance of Logger and creates a new Instance if none exists yet.
func NewSingleton(w io.Writer) *Logger {

	once.Do(func() {
		instance = Logger{
			logger: log.New(w, "", log.Ldate|log.Ltime|log.Lmsgprefix),
		}
	})

	return &instance
}

func (l *Logger) Debug(msg string, v ...interface{}) {
	l.logger.SetPrefix("[DEBUG]: ")
	l.logger.Printf(msg, v...)
}

func (l *Logger) Info(msg string, v ...interface{}) {
	l.logger.SetPrefix("[INFO]: ")
	l.logger.Printf(msg, v...)
}

func (l *Logger) Warn(msg string, v ...interface{}) {
	l.logger.SetPrefix("[WARN]: ")
	l.logger.Printf(msg, v...)
}

func (l *Logger) Error(msg string, v ...interface{}) {
	l.logger.SetPrefix("[ERROR]: ")
	l.logger.Printf(msg, v...)
}

func (l *Logger) Fatal(msg string, v ...interface{}) {
	l.logger.SetPrefix("[FATAL]: ")
	l.logger.Printf(msg, v...)
}
