package protoio

import (
	"encoding/binary"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
)

type Encoder[T proto.Message] struct {
	w io.Writer

	buf []byte
}

func NewEncoder[T proto.Message](w io.Writer) *Encoder[T] {
	return &Encoder[T]{
		w:   w,
		buf: make([]byte, binary.MaxVarintLen64),
	}
}

func (e *Encoder[T]) Encode(msg T) error {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshalling proto: %w", err)
	}

	messageLength := len(bytes)

	nBytesLen := binary.PutUvarint(e.buf, uint64(messageLength))
	_, err = e.w.Write(e.buf[:nBytesLen])
	if err != nil {
		return fmt.Errorf("writing message length: %w", err)
	}

	_, err = e.w.Write(bytes)
	if err != nil {
		return fmt.Errorf("writing message bytes: %w", err)
	}

	return nil
}
