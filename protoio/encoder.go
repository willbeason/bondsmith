package protoio

import (
	"encoding/binary"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
)

type ProtoEncoder[T proto.Message] struct {
	w io.Writer

	buf []byte
}

func NewProtoEncoder[T proto.Message](w io.Writer) *ProtoEncoder[T] {
	return &ProtoEncoder[T]{
		w:   w,
		buf: make([]byte, binary.MaxVarintLen64),
	}
}

func (e *ProtoEncoder[T]) Encode(msg T) error {
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
