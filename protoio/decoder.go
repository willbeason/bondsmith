package protoio

import (
	"encoding/binary"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
)

type Reader interface {
	io.Reader
	io.ByteReader
}

type ProtoDecoder[T proto.Message] struct {
	r Reader

	buf []byte
}

func NewProtoDecoder[T proto.Message](r Reader) *ProtoDecoder[T] {
	return &ProtoDecoder[T]{r: r}
}

func (d *ProtoDecoder[T]) Decode(msg T) error {
	messageLength, err := binary.ReadUvarint(d.r)
	if err != nil {
		return fmt.Errorf("reading message length: %w", err)
	}

	if len(d.buf) < int(messageLength) {
		d.buf = make([]byte, messageLength)
	}

	_, err = io.ReadFull(d.r, d.buf[:messageLength])
	if err != nil {
		return fmt.Errorf("reading message bytes: %w", err)
	}

	err = proto.Unmarshal(d.buf[:messageLength], msg)
	if err != nil {
		return fmt.Errorf("unmarshaling message: %w", err)
	}

	return nil
}
