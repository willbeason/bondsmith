package jsonio

import (
	"encoding/json"
	"io"
	"iter"
)

// Writer writes objects of type T to a Writer.
type Writer[T any] struct {
	w io.Writer

	seq iter.Seq[T]
}

func NewJsonWriter[T any](w io.Writer, seq iter.Seq[T]) *Writer[T] {
	return &Writer[T]{
		w:   w,
		seq: seq,
	}
}

// Write consumes the sequence of objects, encoding them to the Writer.
func (w *Writer[T]) Write() error {
	encoder := json.NewEncoder(w.w)

	for obj := range w.seq {
		err := encoder.Encode(obj)
		if err != nil {
			return err
		}
	}

	return nil
}
