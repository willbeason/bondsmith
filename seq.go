package bondsmith

import "iter"

func Chan2Seq[T any](c chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			v, ok := <-c
			if !ok {
				return
			}

			if !yield(v) {
				return
			}
		}
	}
}
