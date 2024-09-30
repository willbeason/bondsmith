package protoio

import (
	"encoding/binary"
	"fmt"
	"github.com/willbeason/bondsmith"
	"google.golang.org/protobuf/proto"
	"io"
)

type ProtoDecoder[T proto.Message] struct {
	r bondsmith.Reader

	buf []byte
}

func NewProtoDecoder[T proto.Message](r bondsmith.Reader) *ProtoDecoder[T] {
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
