package jsonio

import (
	"encoding/json"
	"io"
	"iter"
)

// Reader reads objects of type T from a Reader.
type Reader[T any] struct {
	r io.Reader

	newValue func() T
}

func NewJsonReader[T any](r io.Reader, newValue func() T) *Reader[T] {
	return &Reader[T]{
		r:        r,
		newValue: newValue,
	}
}

// Read returns an iter.Seq which sequentially decodes Json objects from the reader.
func (r *Reader[T]) Read() iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		decoder := json.NewDecoder(r.r)

		for {
			v := r.newValue()

			err := decoder.Decode(v)
			if !yield(v, err) {
				return
			}
		}
	}
}
