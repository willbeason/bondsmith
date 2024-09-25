package bondsmith_io

import (
	"io"
)

// CountReader counts the number of bytes read from the underlying Reader.
type CountReader struct {
	io.Reader

	Count int64
}

func NewCountReader(r io.Reader) *CountReader {
	return &CountReader{
		Reader: r,
	}
}

func (r *CountReader) Read(p []byte) (n int, err error) {
	read, err := r.Reader.Read(p)
	r.Count += int64(read)

	return read, err
}
