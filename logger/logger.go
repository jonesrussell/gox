package logger

import (
	"io"
	"log"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

// LoggerInterface is an interface for different types of debugging backends
type LoggerInterface interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, err error, args ...any)
}

// Logger is a structured logger that logs to both file and stdout
type Logger struct {
	logger *slog.Logger
}

// Ensure Logger implements LoggerInterface
var _ LoggerInterface = &Logger{}

// NewLogger creates a new instance of Logger
func NewLogger(logFilePath string) (LoggerInterface, error) {
	// Open the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return nil, err
	}

	// Create a JSON handler for the file with immediate flushing
	fileHandler := slog.NewJSONHandler(struct {
		io.Writer
		Sync func() error
	}{
		Writer: file,
		Sync:   file.Sync,
	}, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Create a text handler for stdout
	stdoutHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Combine both handlers using slog-multi
	multiHandler := slogmulti.Fanout(fileHandler, stdoutHandler)

	// Create the logger
	logger := slog.New(multiHandler)

	// Test log to file
	logger.Info("Logger initialized", "file", logFilePath)

	return &Logger{logger: logger}, nil
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info logs an info message
func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs an error message
func (l *Logger) Error(msg string, err error, args ...any) {
	l.logger.Error(msg, append(args, "error", err)...)
}
