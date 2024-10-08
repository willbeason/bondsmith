package protoio

import (
	"github.com/willbeason/bondsmith"
	"google.golang.org/protobuf/proto"
	"iter"
)

// Reader reads objects of type T from a Reader from a stream of protos.
type Reader[T proto.Message] struct {
	r bondsmith.Reader

	newValue func() T
}

func NewReader[T proto.Message](r bondsmith.Reader) *Reader[T] {
	return &Reader[T]{r: r}
}

func (r *Reader[T]) Read() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		decoder := NewDecoder[T](r.r)

		for {
			v := r.newValue()

			err := decoder.Decode(v)
			if !yield(v, err) {
				return
			}
		}
	}
}
