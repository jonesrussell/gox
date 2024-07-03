package debug

import (
	"io"
	"log"
	"os"
)

// Debugger is an interface for different type of debugging backends
type Debugger interface {
	// Debug receives a new debug message.
	Debug(message string)
}

// LogDebugger is the simplest debugger which prints log messages to the STDERR
type LogDebugger struct {
	// Output is the log destination, anything can be used which implements them
	// io.Writer interface. Leave it blank to use STDERR
	Output io.Writer
	logger *log.Logger
}

// NewLogDebugger creates a new instance of LogDebugger
func NewLogDebugger() *LogDebugger {
	debugger := &LogDebugger{}
	debugger.Init()
	return debugger
}

// Init initializes the LogDebugger
func (l *LogDebugger) Init() error {
	if l.Output == nil {
		l.Output = os.Stderr
	}
	l.logger = log.New(l.Output, "", 0)
	return nil
}

// Debug receives a debug message and prints it to STDERR
func (l *LogDebugger) Debug(message string) {
	l.logger.Println(message)
}
