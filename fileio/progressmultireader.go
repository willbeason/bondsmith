package fileio

import (
	"fmt"
	"os"
)

type ProgressMultiReader struct {
	mr *MultiReader

	totalBytesRead int64
	totalBytes     int64
}

func NewProgressMultiReader(filepaths []string) (*ProgressMultiReader, error) {
	totalBytes := int64(0)
	for _, filepath := range filepaths {
		stat, err := os.Stat(filepath)
		if err != nil {
			return nil, fmt.Errorf("statting file: %w", err)
		}

		totalBytes += stat.Size()
	}

	return &ProgressMultiReader{
		mr:         NewMultiFileReader(filepaths),
		totalBytes: totalBytes,
	}, nil
}

func (mr *ProgressMultiReader) TotalBytesRead() int64 {
	return mr.totalBytesRead
}

func (mr *ProgressMultiReader) TotalBytes() int64 {
	return mr.totalBytes
}

func (mr *ProgressMultiReader) Close() error {
	return mr.mr.Close()
}

func (mr *ProgressMultiReader) Read(p []byte) (int, error) {
	n, err := mr.mr.Read(p)
	mr.totalBytesRead += int64(n)
	return n, err
}

func (mr *ProgressMultiReader) ReadByte() (byte, error) {
	b, err := mr.mr.ReadByte()
	mr.totalBytesRead++
	return b, err
}
