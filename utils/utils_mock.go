package utils

// MockFileReader is a mock implementation of the FileReader interface
type MockFileReader struct{}

// ReadFile returns a predefined byte slice
func (fr MockFileReader) ReadFile(filename string) ([]byte, error) {
	return []byte("<html><head><title>Mock Title</title></head><body>Mock Body</body></html>"), nil
}
