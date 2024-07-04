package logger

import (
	"io"
	"log"
	"os"
)

// LoggerInterface is an interface for different types of debugging backends
type LoggerInterface interface {
	Init() error
	Log(message string)
	Debug(message string)
	Error(message string, err error)
}

// Logger is the simplest logger which prints log messages to the specified log file
type Logger struct {
	Output   io.Writer
	instance *log.Logger
}

// Ensure Logger implements LoggerInterface
var _ LoggerInterface = &Logger{}

// NewLogger creates a new instance of Logger
func NewLogger(logFilePath string) LoggerInterface {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	logger := &Logger{
		Output:   file,
		instance: log.New(file, "", 0),
	}

	err = logger.Init()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	return logger
}

// Init initializes the Logger
func (l *Logger) Init() error {
	if l.Output == nil {
		l.Output = os.Stderr
	}
	l.instance = log.New(l.Output, "", 0)
	return nil
}

// Log prints the log message to the log file
func (l *Logger) Log(message string) {
	l.instance.Println(message)
}

// Debug prints the debug message to the log file
func (l *Logger) Debug(message string) {
	l.instance.Println(message)
}

// Error logs an error message with the given error
func (l *Logger) Error(message string, err error) {
	l.instance.Printf("%s: %v\n", message, err)
}
