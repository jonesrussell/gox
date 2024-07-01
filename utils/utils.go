package utils

import "os"

// FileReader is an interface that has a ReadFile method
type FileReader interface {
	ReadFile(filename string) ([]byte, error)
}

// OSFileReader is a struct that will implement the FileReader interface
type OSFileReader struct{}

// ReadFile is the OSFileReader's implementation of the ReadFile method
func (fr OSFileReader) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}