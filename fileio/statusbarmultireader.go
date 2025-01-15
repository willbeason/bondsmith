package fileio

import (
	"fmt"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	"golang.org/x/crypto/ssh/terminal"
)

type StatusBarMultiReader struct {
	mr                     *ProgressMultiReader
	p                      *mpb.Progress
	bar                    *mpb.Bar
	previousTotalBytesRead int64
}

func NewStatusBarMultiReader(filepaths []string) (*StatusBarMultiReader, error) {
	mr, err := NewProgressMultiReader(filepaths)
	if err != nil {
		return nil, fmt.Errorf("creating progress multireader: %w", err)
	}

	width, _, err := terminal.GetSize(0)
	if err != nil {
		return nil, fmt.Errorf("getting terminal width: %w", err)
	}

	p := mpb.New(mpb.WithWidth(width))
	bar := p.AddBar(mr.totalBytes,
		mpb.AppendDecorators(decor.AverageETA(decor.ET_STYLE_GO)),
		mpb.BarRemoveOnComplete(),
	)

	return &StatusBarMultiReader{
		mr:  mr,
		p:   p,
		bar: bar,
	}, nil
}

func (mr *StatusBarMultiReader) UpdateStatusBar() {
	incrementBytesRead := mr.mr.TotalBytesRead() - mr.previousTotalBytesRead
	mr.previousTotalBytesRead = mr.mr.TotalBytesRead()
	mr.bar.IncrBy(int(incrementBytesRead))
}

func (mr *StatusBarMultiReader) Close() error {
	return mr.mr.Close()
}

func (mr *StatusBarMultiReader) Read(p []byte) (int, error) {
	n, err := mr.mr.Read(p)
	return n, err
}

func (mr *StatusBarMultiReader) ReadByte() (byte, error) {
	b, err := mr.mr.ReadByte()
	return b, err
}
