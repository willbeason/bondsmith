package bondsmith

import (
	"io"
)

// CountReader counts the number of bytes read from the underlying Reader.
type CountReader struct {
	io.Reader

	count int64
}

func NewCountReader(r io.Reader) *CountReader {
	return &CountReader{
		Reader: r,
	}
}

func (r *CountReader) Read(p []byte) (int, error) {
	read, err := r.Reader.Read(p)
	r.count += int64(read)

	return read, err
}

func (r *CountReader) Count() int64 {
	return r.count
}
