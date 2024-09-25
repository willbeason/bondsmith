package bondsmith_io

import (
	"google.golang.org/protobuf/proto"
	"iter"
)

type ProtoReader[T proto.Message] struct {
	r Reader

	newValue func() T
}

func NewProtoReader[T proto.Message](r Reader) *ProtoReader[T] {
	return &ProtoReader[T]{r: r}
}

type ProtoReaderOpt[T proto.Message] func(*ProtoReader[T])

func (r *ProtoReader[T]) Read() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		decoder := NewProtoDecoder[T](r.r)

		for {
			v := r.newValue()

			err := decoder.Decode(v)
			if !yield(v, err) {
				return
			}
		}
	}
}
