package protoio

import (
	"google.golang.org/protobuf/proto"
	"io"
	"iter"
)

type Writer[T proto.Message] struct {
	w io.Writer

	seq iter.Seq[T]
}

func NewWriter[T proto.Message](w io.Writer, seq iter.Seq[T]) *Writer[T] {
	return &Writer[T]{
		w:   w,
		seq: seq,
	}
}

func (w *Writer[T]) Write() error {
	encoder := NewEncoder[T](w.w)

	for obj := range w.seq {
		err := encoder.Encode(obj)
		if err != nil {
			return err
		}
	}

	return nil
}
