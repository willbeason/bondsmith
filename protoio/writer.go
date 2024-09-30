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

func NewProtoWriter[T proto.Message](w io.Writer, seq iter.Seq[T], opts ...ProtoWriterOpt[T]) *Writer[T] {
	result := &Writer[T]{
		w:   w,
		seq: seq,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

type ProtoWriterOpt[T proto.Message] func(*Writer[T])

func (w *Writer[T]) Write() error {
	encoder := NewProtoEncoder[T](w.w)

	for obj := range w.seq {
		err := encoder.Encode(obj)
		if err != nil {
			return err
		}
	}

	return nil
}
