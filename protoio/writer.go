package protoio

import (
	"google.golang.org/protobuf/proto"
	"io"
	"iter"
)

type ProtoWriter[T proto.Message] struct {
	w io.Writer

	seq iter.Seq[T]
}

func NewProtoWriter[T proto.Message](w io.Writer, seq iter.Seq[T], opts ...ProtoWriterOpt[T]) *ProtoWriter[T] {
	result := &ProtoWriter[T]{
		w:   w,
		seq: seq,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

type ProtoWriterOpt[T proto.Message] func(*ProtoWriter[T])

func (w *ProtoWriter[T]) Write() error {
	encoder := NewProtoEncoder[T](w.w)

	for obj := range w.seq {
		err := encoder.Encode(obj)
		if err != nil {
			return err
		}
	}

	return nil
}
