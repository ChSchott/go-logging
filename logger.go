package logger

import (
	"io"
	"log"
)

type ILogger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

type Logger struct {
	logger *log.Logger
}

func New(w io.Writer) *Logger {
	return &Logger{
		logger: log.New(w, "", log.Ldate|log.Ltime|log.Lmsgprefix),
	}
}

func (l *Logger) Debug(msg string) {
	l.logger.SetPrefix("[DEBUG]: ")
	l.logger.Println(msg)
}

func (l *Logger) Info(msg string) {
	l.logger.SetPrefix("[INFO]: ")
	l.logger.Println(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.SetPrefix("[WARN]: ")
	l.logger.Println(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.SetPrefix("[ERROR]: ")
	l.logger.Println(msg)
}

func (l *Logger) Fatal(msg string) {
	l.logger.SetPrefix("[FATAL]: ")
	l.logger.Println(msg)
}
