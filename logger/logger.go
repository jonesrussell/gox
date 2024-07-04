package logger

import (
	"io"
	"log"
	"os"
)

// LoggerInterface is an interface for different type of debugging backends
type LoggerInterface interface {
	Init() error
	Log(message string)
	// Debug receives a new debug message.
	Debug(message string)
}

// Logger is the simplest logger which prints log messages to the STDERR
type Logger struct {
	// Output is the log destination, anything can be used which implements them
	// io.Writer interface. Leave it blank to use STDERR
	Output    io.Writer
	logLogger *log.Logger
}

// Ensure menuImpl implements MenuInterface
var _ LoggerInterface = &Logger{}

// NewLogger creates a new instance of Logger
func NewLogger() LoggerInterface {
	logger := &Logger{}
	err := logger.Init()
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
	l.logLogger = log.New(l.Output, "", 0)
	return nil
}

func (l *Logger) Log(message string) {
	log.Println(message)
}

// Debug receives a debug message and prints it to STDERR
func (l *Logger) Debug(message string) {
	l.logLogger.Println(message)
}
