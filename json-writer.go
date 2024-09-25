package bondsmith_io

import (
	"encoding/json"
	"fmt"
	"io"
	"iter"
)

// JsonWriter writes objects of type T to a Writer.
type JsonWriter[T any] struct {
	w io.Writer

	seq iter.Seq[T]

	reporter ErrorReporter
}

func NewJsonWriter[T any](w io.Writer, seq iter.Seq[T], opts ...JsonWriterOpt[T]) *JsonWriter[T] {
	result := &JsonWriter[T]{
		w:   w,
		seq: seq,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

type JsonWriterOpt[T any] func(*JsonWriter[T])

func JsonWriterErrorReporter[T any](reporter ErrorReporter) JsonWriterOpt[T] {
	return func(w *JsonWriter[T]) {
		w.reporter = reporter
	}
}

// Write consumes the sequence of objects, encoding them to the Writer.
// Errors are suppressed unless an error reporting option is passed at creation time.
func (w *JsonWriter[T]) Write() {
	encoder := json.NewEncoder(w.w)

	for obj := range w.seq {
		err := encoder.Encode(obj)
		if err != nil {
			if w.reporter != nil {
				err = fmt.Errorf("encoding json: %w", err)
				w.reporter.Report(err)
			}
			return
		}
	}
}
