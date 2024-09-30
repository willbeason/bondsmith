package bondsmith

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// MultiFileReader constructs a Reader that is the logical concatenation of the
// readers of the passed filepaths. Files are read sequentially. Once all readers
// have returned EOF, Read will return EOF. If any reader returns any other error,
// Read returns that error.
type MultiFileReader struct {
	filepaths []string

	closer io.Closer
	reader *bufio.Reader
}

var _ Reader = &MultiFileReader{}

func NewMultiFileReader(filepaths []string) *MultiFileReader {
	return &MultiFileReader{filepaths: filepaths}
}

func (mr *MultiFileReader) getReader() (*bufio.Reader, error) {
	if mr.reader != nil {
		return mr.reader, nil
	}

	if len(mr.filepaths) == 0 {
		return nil, io.EOF
	}

	fileReader, err := os.Open(mr.filepaths[0])
	if err != nil {
		return nil, err
	}
	mr.closer = fileReader
	mr.reader = bufio.NewReader(fileReader)

	mr.filepaths = mr.filepaths[1:]

	return mr.reader, nil
}

func (mr *MultiFileReader) Read(p []byte) (int, error) {
	reader, err := mr.getReader()
	if err != nil {
		return 0, err
	}

	n, err := reader.Read(p)
	if err == io.EOF {
		err = mr.closer.Close()
		if err != nil {
			return 0, fmt.Errorf("closing file: %w", err)
		}

		mr.reader = nil
		return mr.Read(p)
	}

	return n, err

}

func (mr *MultiFileReader) ReadByte() (byte, error) {
	reader, err := mr.getReader()
	if err != nil {
		return 0, err
	}

	b, err := reader.ReadByte()
	if err == io.EOF {
		err = mr.closer.Close()
		if err != nil {
			return 0, fmt.Errorf("closing file: %w", err)
		}

		mr.reader = nil
		return mr.ReadByte()
	}

	return b, err
}
