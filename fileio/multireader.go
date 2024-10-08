package fileio

import (
	"bufio"
	"fmt"
	"github.com/willbeason/bondsmith"
	"io"
	"os"
)

// MultiReader constructs a Reader that is the logical concatenation of the
// readers of the passed filepaths. Files are read sequentially. Each file is
// closed as it returns EOF, and at most one file is open at any time. Once all
// readers have returned EOF, Read will return EOF. If any reader returns any
// other error, Read returns that error.
type MultiReader struct {
	filepaths []string

	closer io.Closer
	reader *bufio.Reader
}

var _ bondsmith.Reader = &MultiReader{}

func NewMultiFileReader(filepaths []string) *MultiReader {
	return &MultiReader{filepaths: filepaths}
}

func (mr *MultiReader) getReader() (*bufio.Reader, error) {
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

// Close closes the currently-opened file if one is currently opened.
// Does nothing and returns no error if no files are currently opened.
func (mr *MultiReader) Close() error {
	if mr.closer == nil {
		return nil
	}

	err := mr.closer.Close()
	if err != nil {
		return fmt.Errorf("closing file: %w", err)
	}

	mr.closer = nil
	mr.reader = nil

	return nil
}

func (mr *MultiReader) Read(p []byte) (int, error) {
	reader, err := mr.getReader()
	if err != nil {
		return 0, err
	}

	n, err := reader.Read(p)
	if err == io.EOF {
		err = mr.Close()
		if err != nil {
			return 0, err
		}

		return mr.Read(p)
	}

	return n, err

}

func (mr *MultiReader) ReadByte() (byte, error) {
	reader, err := mr.getReader()
	if err != nil {
		return 0, err
	}

	b, err := reader.ReadByte()
	if err == io.EOF {
		err = mr.Close()
		if err != nil {
			return 0, err
		}

		return mr.ReadByte()
	}

	return b, err
}
