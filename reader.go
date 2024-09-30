package bondsmith

import "io"

type Reader interface {
	io.Reader
	io.ByteReader
}
