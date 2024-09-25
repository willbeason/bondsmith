package bondsmith_io

import (
	"encoding/json"
	"io"
	"iter"
)

// JsonReader reads objects of type T from a Reader.
type JsonReader[T any] struct {
	r io.Reader

	newValue func() T
}

func NewJsonReader[T any](r io.Reader, newValue func() T) *JsonReader[T] {
	return &JsonReader[T]{
		r:        r,
		newValue: newValue,
	}
}

// Read returns an iter.Seq which sequentially decodes Json objects from the reader.
// Errors are suppressed unless an error reporter is passed at creation time.
func (r *JsonReader[T]) Read() iter.Seq2[T, error] {
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
