package fileio

// BytesReadReader tracks the number of bytes read from the passed Reader.
type BytesReadReader struct {
	Reader
	bytesRead int64
}

func NewBytesReadMultiReader(reader Reader) *BytesReadReader {
	return &BytesReadReader{
		Reader: reader,
	}
}

func (mr *BytesReadReader) BytesRead() int64 {
	return mr.bytesRead
}

func (mr *BytesReadReader) Read(p []byte) (int, error) {
	n, err := mr.Reader.Read(p)
	mr.bytesRead += int64(n)
	return n, err
}

func (mr *BytesReadReader) ReadByte() (byte, error) {
	b, err := mr.Reader.ReadByte()
	mr.bytesRead++
	return b, err
}
