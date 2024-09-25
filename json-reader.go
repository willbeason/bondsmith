package bondsmith_io

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
)

// JsonReader reads objects of type T from a Reader.
type JsonReader[T any] struct {
	r io.Reader

	newValue func() T

	reporter ErrorReporter
}

func NewJsonReader[T any](r io.Reader, newValue func() T, opts ...JsonReaderOpt[T]) *JsonReader[T] {
	result := &JsonReader[T]{
		r:        r,
		newValue: newValue,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

type JsonReaderOpt[T any] func(reader *JsonReader[T])

func JsonReaderErrorReporter[T any](reporter ErrorReporter) JsonReaderOpt[T] {
	return func(reader *JsonReader[T]) {
		reader.reporter = reporter
	}
}

// Read returns an iter.Seq which sequentially decodes Json objects from the reader.
// Errors are suppressed unless an error reporter is passed at creation time.
func (r *JsonReader[T]) Read() iter.Seq[T] {
	return func(yield func(T) bool) {
		decoder := json.NewDecoder(r.r)

		for {
			v := r.newValue()
			err := decoder.Decode(v)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					if r.reporter != nil {
						err = fmt.Errorf("decoding json: %w", err)
						r.reporter.Report(err)
					}
				}

				return
			}

			if !yield(v) {
				return
			}
		}
	}
}
