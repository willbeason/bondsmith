package bondsmith_io

import (
	"encoding/json"
	"io"
	"iter"
)

// JsonWriter writes objects of type T to a Writer.
type JsonWriter[T any] struct {
	w io.Writer

	seq iter.Seq[T]
}

func NewJsonWriter[T any](w io.Writer, seq iter.Seq[T]) *JsonWriter[T] {
	return &JsonWriter[T]{
		w:   w,
		seq: seq,
	}
}

// Write consumes the sequence of objects, encoding them to the Writer.
// Errors are suppressed unless an error reporting option is passed at creation time.
func (w *JsonWriter[T]) Write() error {
	encoder := json.NewEncoder(w.w)

	for obj := range w.seq {
		err := encoder.Encode(obj)
		if err != nil {
			return err
		}
	}

	return nil
}
