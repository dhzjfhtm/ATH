package record

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	fpLog, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Println("INFO", msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Println("Error ", msg)
}
